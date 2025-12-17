package repository

import (
	"context"
	"errors"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	// Inser a new user
	CreatNewUser(ctx context.Context, user *model.User) (*model.User, error)
	VerifyMail(ctx context.Context, mail string) (*model.User, error)
}

type userRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

// constructor to get the userRepository
func GetUserRepo(l *slog.Logger, db *gorm.DB) UserRepository {
	return &userRepository{
		db:     db,
		logger: l,
	}
}

// Insert a new user -> signup
func (ur *userRepository) CreatNewUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := ur.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// validate email
func (ur *userRepository) VerifyMail(ctx context.Context, email string) (*model.User, error) {

	var user model.User

	err := ur.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		// actual db error
		ur.logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}
