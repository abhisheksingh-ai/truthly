package model

import (
	"time"

	"github.com/google/uuid"
)

type Description struct {
	// fk
	DescriptionId string `gorm:"column:DescriptionId; primaryKey"`
	ImageId       string `gorm:"column:ImageId"`
	UserId        string `gorm:"column:UserId"`

	// These details user will post along with image
	Description string `gorm:"column:Description"`
	Country     string `gorm:"column:Country"`
	State       string `gorm:"column:State"`
	City        string `gorm:"column:City"`

	// dates
	CreatedAt time.Time `gorm:"column:CreatedAt; autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt; autoUpdateTime"`
}

func (Description) TableName() string {
	return "Descriptions"
}

func (d *Description) BeforeCreate() (err error) {
	if d.DescriptionId == "" {
		d.DescriptionId = uuid.New().String()
	}
	return
}
