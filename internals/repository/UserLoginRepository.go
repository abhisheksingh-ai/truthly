package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type UserLoginRepository interface{}

// for UserLogin Table
type userLoginRepo struct {
	logger *slog.Logger
	Db     *gorm.DB
}

func GetNewUserLoginRepo(logger *slog.Logger, Db *gorm.DB) UserLoginRepository {
	return &userLoginRepo{
		logger: logger,
		Db:     Db,
	}
}
