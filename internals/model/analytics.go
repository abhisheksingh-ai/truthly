package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Analytic struct {
	AnalyticId    string `gorm:"column:AnalyticId; primaryKey"`
	ImageId       string `gorm:"column:ImageId"`
	DescriptionId string `gorm:"column:DescriptionId"`
	UserId        string `gorm:"column:UserId"`

	Like    int `gorm:"column:Like"`
	Share   int `gorm:"column:Share"`
	Comment int `gorm:"column:Comment"`

	CreatedAt time.Time `gorm:"column:CreatedAt; autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt; autoUpdateTime"`
}

func (Analytic) TableName() string {
	return "Analytic"
}

func (a *Analytic) BeforeCreate(tx *gorm.DB) (err error) {
	if a.AnalyticId == "" {
		a.AnalyticId = uuid.New().String()
	}
	return
}
