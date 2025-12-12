package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type ImageRepository interface {
	InsertNewImage(ctx context.Context, data *model.Image) (*model.Image, error)
}

type imageRepository struct {
	Db     *gorm.DB
	logger *slog.Logger
}

// constructor
func GetImageRepo(Db *gorm.DB, logger *slog.Logger) ImageRepository {
	return &imageRepository{
		Db:     Db,
		logger: logger,
	}
}

func (i *imageRepository) InsertNewImage(ctx context.Context, data *model.Image) (*model.Image, error) {
	if err := i.Db.WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
