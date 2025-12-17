package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type UserSessionRepository interface{}

// for userSessionRepo
type userSessionRepo struct {
	logger *slog.Logger
	Db     *gorm.DB
}

func GetNewUserSessionRepo(logger *slog.Logger, Db *gorm.DB) UserSessionRepository {
	return &userSessionRepo{
		logger: logger,
		Db:     Db,
	}
}
