package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type DescriptionRepository interface {
	InsertDescription(ctx context.Context, description *model.Description) (*model.Description, error)
}

type descriptionRepository struct {
	Db     *gorm.DB
	logger *slog.Logger
}

// constructor
func GetDescriptionRepository(Db *gorm.DB, logger *slog.Logger) DescriptionRepository {
	return &descriptionRepository{
		Db:     Db,
		logger: logger,
	}
}

func (d *descriptionRepository) InsertDescription(ctx context.Context, description *model.Description) (*model.Description, error) {
	if err := d.Db.WithContext(ctx).Create(description).Error; err != nil {
		return nil, err
	}
	return description, nil
}
