package service

import (
	"context"
	"log/slog"
	"truthly/internals/dto"
	"truthly/internals/repository"
)

type PostService interface {
	UploadPost(ctx context.Context, imageUrl string) (*dto.ResponseDto, error)
}

type postService struct {
	logger          *slog.Logger
	analyticsRepo   *repository.AnalyticRepository
	commentRepo     *repository.CommentRepository
	descriptionRepo *repository.DescriptionRepository
	imageRepo       *repository.ImageRepository
}

func GetPostService(
	logger *slog.Logger,
	analyticsRepo *repository.AnalyticRepository,
	commentRepo *repository.CommentRepository,
	descriptionRepo *repository.DescriptionRepository,
	imageRepo *repository.ImageRepository,
) PostService {
	return &postService{
		logger:          logger,
		analyticsRepo:   analyticsRepo,
		commentRepo:     commentRepo,
		descriptionRepo: descriptionRepo,
		imageRepo:       imageRepo,
	}
}

func (s *postService) UploadPost(ctx context.Context, imageUrl string) (*dto.ResponseDto, error) {
	var result dto.ResponseDto
	return &result, nil
}
