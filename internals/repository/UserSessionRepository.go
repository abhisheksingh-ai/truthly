package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type UserSessionRepository interface {
	CreateNewSession(
		ctx context.Context,
		userSession *model.UserSession,
	) error
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
