package service

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	Client     *s3.Client
	BucketName string
	logger     *slog.Logger
}

func NewS3Uploader(bucket string, logger *slog.Logger) (*S3Uploader, error) {

	region := os.Getenv("AWS_REGION")
	if region == "" {
		return nil, fmt.Errorf("AWS_REGION is not set")
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)

	if err != nil {
		return nil, err
	}

	return &S3Uploader{
		Client:     s3.NewFromConfig(cfg),
		BucketName: bucket,
		logger:     logger,
	}, nil
}

func (s *S3Uploader) UploadImage(fileHeader *multipart.FileHeader, fileName string) (string, error) {
	// 1. Open the file
	file, err := fileHeader.Open()
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	defer file.Close()

	// 2. Upload to S3
	_, err = s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.BucketName,
		Key:    &fileName,
		Body:   file,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	// 3. Generate the image URL
	// this is cloudfront domain
	url := fmt.Sprintf("https://d316g795lwuaau.cloudfront.net/%s", fileName)
	s.logger.Info("Image is uploaded on s3 bucket", "imageUrl", url)

	return url, nil
}
