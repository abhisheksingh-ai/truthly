package service

import (
	"context"
	"log/slog"
	"mime/multipart"
	"truthly/internals/dto"
	"truthly/internals/repository"
)

type PostService interface {
	UploadPost(ctx context.Context, file multipart.File) (*dto.ResponseDto, error)
}

type postService struct {
	logger          *slog.Logger
	analyticsRepo   *repository.AnalyticRepository
	commentRepo     *repository.CommentRepository
	descriptionRepo *repository.DescriptionRepository
	imageRepo       *repository.ImageRepository
	s3Uploader      *S3Uploader
}

func GetPostService(
	logger *slog.Logger,
	analyticsRepo *repository.AnalyticRepository,
	commentRepo *repository.CommentRepository,
	descriptionRepo *repository.DescriptionRepository,
	imageRepo *repository.ImageRepository,
	s3Uploader *S3Uploader,
) PostService {
	return &postService{
		logger:          logger,
		analyticsRepo:   analyticsRepo,
		commentRepo:     commentRepo,
		descriptionRepo: descriptionRepo,
		imageRepo:       imageRepo,
		s3Uploader:      s3Uploader,
	}
}

func (s *postService) UploadPost(ctx context.Context, file multipart.File) (*dto.ResponseDto, error) {
	var result dto.ResponseDto

	//1. Upload the image on s3 bucket and get url

	//2. Insert row in Image table

	// 3. Insert the description row in table like city, state etc

	// 4. Analytic like, share, comment initially set 0

	// 5. Comment initally nothing

	return &result, nil
}
