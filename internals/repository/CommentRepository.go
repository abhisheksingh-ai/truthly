package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment *model.Commemts) (*model.Commemts, error)
	GetRecentComment(ctx context.Context, imageIds []string) ([]commentRow, error)
}

// For comment row data
type commentRow struct {
	CommentId string
	Comment   string
	ImageId   string
	UserId    string
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

// Get one comment for each imageId
func (c *commentRepository) GetRecentComment(ctx context.Context, imageIds []string) ([]commentRow, error) {
	var rows []commentRow

	query := `
		SELECT CommentId, Comment, ImageId, UserId
		FROM (
		  SELECT
		    CommentId,
		    Comment,
		    ImageId,
			UserId,
		    ROW_NUMBER() OVER (
		      PARTITION BY ImageId
		      ORDER BY CreatedAt DESC
		    ) AS rn
		  FROM Comments
		  WHERE ImageId IN ?
		) ranked
		WHERE rn <= 1
	`

	err := c.Db.WithContext(ctx).
		Raw(query, imageIds).
		Scan(&rows).Error

	return rows, err
}
