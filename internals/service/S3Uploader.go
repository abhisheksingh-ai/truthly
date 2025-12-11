package service

import (
	"context"

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
