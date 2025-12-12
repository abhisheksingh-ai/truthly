package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type AnalyticRepository interface {
	InsertAnalytics(ctx context.Context, analytic *model.Analytic) (*model.Analytic, error)
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

func (a *analyticRepository) InsertAnalytics(ctx context.Context, analytic *model.Analytic) (*model.Analytic, error) {
	if err := a.Db.WithContext(ctx).Create(analytic).Error; err != nil {
		return nil, err
	}
	return analytic, nil
}
