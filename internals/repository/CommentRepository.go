package repository

import (
	"log/slog"

	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment()
}

type commentRepository struct {
	Db     *gorm.DB
	logger *slog.Logger
}

// constructor
func GetCommentRepository(Db *gorm.DB, logger *slog.Logger) CommentRepository {
	return &commentRepository{
		Db:     Db,
		logger: logger,
	}
}

func (c *commentRepository) InsertComment() {

}
