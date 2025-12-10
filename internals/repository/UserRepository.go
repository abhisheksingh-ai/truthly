package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	// Inser a new user
	CreatNewUser(ctx context.Context, user *model.User) (*model.User, error)
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

// Insert a new user
func (ur *userRepository) CreatNewUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := ur.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
