package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment *model.Commemts) (*model.Commemts, error)
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

func (c *commentRepository) InsertComment(ctx context.Context, comment *model.Commemts) (*model.Commemts, error) {
	if err := c.Db.WithContext(ctx).Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
