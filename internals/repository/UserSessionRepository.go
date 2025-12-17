package repository

import (
	"context"
	"errors"
	"log/slog"
	"time"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type UserSessionRepository interface {
	CreateNewSession(
		ctx context.Context,
		userSession *model.UserSession,
	) error

	ExpireLastActiveSession(ctx context.Context, userId string) error

	GetActiveSession(ctx context.Context, sessionId string) (*model.UserSession, error)
}

// for userSessionRepo
type userSessionRepo struct {
	logger *slog.Logger
	Db     *gorm.DB
}

func GetNewUserSessionRepo(
	logger *slog.Logger,
	Db *gorm.DB,
) UserSessionRepository {
	return &userSessionRepo{
		logger: logger,
		Db:     Db,
	}
}

func (us *userSessionRepo) CreateNewSession(
	ctx context.Context,
	userSession *model.UserSession,
) error {

	err := us.Db.WithContext(ctx).Create(userSession).Error
	if err != nil {
		us.logger.Error(
			"failed to create user session",
			"error", err,
			"user_id", userSession.UserId,
		)
		return err
	}
	return nil
}

func (us *userSessionRepo) ExpireLastActiveSession(
	ctx context.Context,
	userId string,
) error {

	var session model.UserSession

	//  Get last ACTIVE session
	err := us.Db.WithContext(ctx).
		Where("UserId = ? AND Status = ?", userId, "ACTIVE").
		Order("CreatedAt DESC").
		Limit(1).
		First(&session).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			us.logger.Warn(
				"no active session found",
				"userId", userId,
			)
			return nil
		}
		return err
	}

	//  Expire ONLY that session
	err = us.Db.WithContext(ctx).
		Model(&model.UserSession{}).
		Where("id = ?", session.Id).
		Updates(map[string]interface{}{
			"Status":    "EXPIRED",
			"ExpiredAt": time.Now(),
		}).Error

	if err != nil {
		us.logger.Error(
			"failed to expire last session",
			"sessionId", session.Id,
			"error", err,
		)
		return err
	}

	return nil
}

func (us *userSessionRepo) GetActiveSession(ctx context.Context, sessionId string) (*model.UserSession, error) {

	var userSession model.UserSession

	err := us.Db.WithContext(ctx).
		Where("SessionId=? AND Status=? AND ExpiredAt > NOW()", sessionId, "ACTIVE").
		First(&userSession).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Session Id with Active status not found")
		}
		return nil, err
	}

	return &userSession, nil
}
