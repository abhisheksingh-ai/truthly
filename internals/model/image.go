package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	ImageId string `gorm:"column:ImageId;primaryKey"`
	UserId  string `gorm:"column:UserId;not null"`
	// s3 bucket image url
	ImageUrl string `gorm:"column:ImageUrl; not null"`

	CreatedAt time.Time `gorm:"column:CreatedAt; autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt; autoUpdateTime"`
}

func (Image) TableName() string {
	return "Image"
}

func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ImageId == "" {
		i.ImageId = uuid.New().String()
	}
	return
}
