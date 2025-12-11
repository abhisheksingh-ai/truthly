package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type ImageRepository interface {
	InsertNewImage()
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

func (i *imageRepository) InsertNewImage() {

}
