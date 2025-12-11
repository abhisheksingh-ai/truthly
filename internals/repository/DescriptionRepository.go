package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type DescriptionRepository interface {
	InsertDescription()
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

func (c *descriptionRepository) InsertDescription() {

}
