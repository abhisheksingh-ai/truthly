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
	// File related data
	fileHeader := postReq.FileHeader

	// unique file name
	fileName := fmt.Sprintf("uploads/%d-%s", time.Now().Unix(), fileHeader.Filename)

	//1. Upload the image on s3 bucket and get url
	imgUrl, err := s.s3Uploader.UploadImage(fileHeader, fileName)
	if err != nil {
		s.logger.Error(err.Error())
		return &dto.ResponseDto{
			Status:    "failed",
			Message:   err.Error(),
			ResultObj: nil,
		}, err
	}

	//2. Insert row in Image table
	img := &model.Image{
		ImageUrl: imgUrl,
		UserId:   postReq.UserId,
	}

	imgRes, err := s.imageRepo.InsertNewImage(ctx, img)
	if err != nil {
		s.logger.Error(err.Error())
		return &dto.ResponseDto{
			Status:    "failed",
			Message:   err.Error(),
			ResultObj: nil,
		}, err
	}

	// 3. Insert the description row in table like city, state etc

	description := &model.Description{
		ImageId: imgRes.ImageId,
		UserId:  imgRes.UserId,

		Description: postReq.Description,
		Country:     postReq.Country,
		City:        postReq.City,
		State:       postReq.State,
	}

	descRes, err := s.descriptionRepo.InsertDescription(ctx, description)
	if err != nil {
		s.logger.Error(err.Error())
		return &dto.ResponseDto{
			Status:    "failed",
			Message:   err.Error(),
			ResultObj: nil,
		}, err
	}

	// 4. Analytic like, share, comment initially set 0

	analytic := &model.Analytic{
		ImageId:       descRes.ImageId,
		DescriptionId: descRes.DescriptionId,
		UserId:        descRes.UserId,
		// other field will be zero
	}

	analyticRes, err := s.analyticsRepo.InsertAnalytics(ctx, analytic)
	if err != nil {
		s.logger.Error(err.Error())
		return &dto.ResponseDto{
			Status:    "failed",
			Message:   err.Error(),
			ResultObj: nil,
		}, err
	}

	// 5. Comment initally nothing
	comment := &model.Commemts{
		UserId:        analyticRes.UserId,
		ImageId:       analyticRes.ImageId,
		DescriptionId: analyticRes.DescriptionId,
		AnalyticId:    analyticRes.AnalyticId,
	}

	_, err = s.commentRepo.InsertComment(ctx, comment)
	if err != nil {
		s.logger.Error(err.Error())
		return &dto.ResponseDto{
			Status:    "failed",
			Message:   err.Error(),
			ResultObj: nil,
		}, err
	}

	return &dto.ResponseDto{
		Status:  "success",
		Message: "Post created",
		ResultObj: map[string]string{
			"imageUrl": imgRes.ImageUrl,
		},
	}, nil
}
