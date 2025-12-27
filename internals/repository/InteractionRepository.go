package repository

import (
	"context"
	"log/slog"
	"truthly/internals/model"

	"gorm.io/gorm"
)

type InteractionRepository interface {
	// to incrase the like on image
	LikeImage(ctx context.Context, userId, imageId string) error
	AddComment(ctx context.Context, userId, imageId, text string) error
}

type interactionRepository struct {
	logger *slog.Logger
	Db     *gorm.DB
}

func GetNewInteractionRepository(db *gorm.DB, logger *slog.Logger) InteractionRepository {
	return &interactionRepository{
		Db:     db,
		logger: logger,
	}
}

func (r *interactionRepository) LikeImage(ctx context.Context, userId, imageId string) error {
	return r.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//1. find analytics for this image
		var analytics model.Analytic

		err := tx.Where("ImageId = ?", imageId).
			First(&analytics).Error

		if err != nil {
			r.logger.Error("analytics not found", "imageId", imageId)
			return err
		}

		// 2. atomic increment
		err = tx.
			Model(&model.Analytic{}).
			Where("AnalyticId = ?", analytics.AnalyticId).
			UpdateColumn("LikeCount", gorm.Expr("LikeCount + ?", 1)).Error

		if err != nil {
			r.logger.Error("Error in increasing the like count", "error", err.Error())
			return err
		}

		return nil
	})
}

func (r *interactionRepository) AddComment(ctx context.Context, userId, imageId, text string) error {
	return r.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//1. Get Analytics row
		var analytic *model.Analytic

		err := tx.Where("Imageid = ?", imageId).
			First(&analytic).Error

		if err != nil {
			r.logger.Error("Aanalytics row not found", "error", err.Error, "imageId", imageId)
			return err
		}

		// 2. Add comment
		comment := &model.Commemts{
			UserId:  userId,
			ImageId: imageId,

			DescriptionId: analytic.DescriptionId,
			AnalyticId:    analytic.AnalyticId,
			Comment:       text,
		}

		if err := tx.Create(comment).Error; err != nil {
			r.logger.Error("failed to create comment", "error", err)
			return err
		}

		// 3. Increment comment count
		if err = tx.
			Model(&model.Analytic{}).
			Where("AnalyticId = ?", analytic.AnalyticId).
			UpdateColumn("CommentCount", gorm.Expr("CommentCount + ?", 1)).Error; err != nil {
			r.logger.Error("failed to increase comment count", "error", err.Error())
			return err
		}

		return nil
	})
}
