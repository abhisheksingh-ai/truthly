package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Uploader(bucket string) (*S3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, err
	}

	return &S3Uploader{
		Client:     s3.NewFromConfig(cfg),
		BucketName: bucket,
	}, nil
}

func (s *S3Uploader) UploadImage(file multipart.File, fileName string) (string, error) {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.BucketName,
		Key:    &fileName,
		Body:   file,
		ACL:    "public-read",
	})

	if err != nil {
		return "", err
	}

	// generate the image url
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.BucketName, fileName)
	return url, nil
}
