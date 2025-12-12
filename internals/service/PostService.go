package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"truthly/internals/dto"
	"truthly/internals/model"
	"truthly/internals/repository"
)

type PostService interface {
	UploadPost(ctx context.Context, postReq *dto.PostRequestDto) (*dto.ResponseDto, error)
}

type postService struct {
	logger          *slog.Logger
	analyticsRepo   repository.AnalyticRepository
	commentRepo     repository.CommentRepository
	descriptionRepo repository.DescriptionRepository
	imageRepo       repository.ImageRepository
	s3Uploader      *S3Uploader
}

func GetPostService(
	logger *slog.Logger,
	analyticsRepo repository.AnalyticRepository,
	commentRepo repository.CommentRepository,
	descriptionRepo repository.DescriptionRepository,
	imageRepo repository.ImageRepository,
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

func (s *postService) UploadPost(ctx context.Context, postReq *dto.PostRequestDto) (*dto.ResponseDto, error) {
	var result dto.ResponseDto

	// File related data
	file := postReq.File
	fileHeader := postReq.FileHeader
	userId := postReq.UserId

	// unique file name
	fileName := fmt.Sprintf("uploads/%d-%s", time.Now().Unix(), fileHeader.Filename)

	//1. Upload the image on s3 bucket and get url
	imgUrl, err := s.s3Uploader.UploadImage(file, fileName)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	//2. Insert row in Image table
	img := &model.Image{
		ImageUrl: imgUrl,
		UserId:   userId,
	}

	imgRes, err := s.imageRepo.InsertNewImage(ctx, img)
	// 3. Insert the description row in table like city, state etc

	// 4. Analytic like, share, comment initially set 0

	// 5. Comment initally nothing

	return &result, nil
}
