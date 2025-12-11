package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type AnalyticRepository interface {
	InsertAnalytics()
}

type analyticRepository struct {
	Db     *gorm.DB
	logger *slog.Logger
}

// constructor
func GetAnalyticRepository(Db *gorm.DB, logger *slog.Logger) AnalyticRepository {
	return &analyticRepository{
		Db:     Db,
		logger: logger,
	}
}

func (a *analyticRepository) InsertAnalytics() {

}
