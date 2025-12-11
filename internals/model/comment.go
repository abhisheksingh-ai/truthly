package model

import "github.com/google/uuid"

type Comment struct {
	UserName string
	Value    string
}

type Commemts struct {
	CommentId     string `gorm:"column:CommentId; primaryKey"`
	UserId        string `gorm:"column:UserId"`
	ImageId       string `gorm:"column:ImageId"`
	DescriptionId string `gorm:"column:DescriptionId"`
	AnalyticId    string `gorm:"column:Analyticid"`

	AllComments []Comment `gorm:"column:Comments"`
}

func (Commemts) TableName() string {
	return "Commemts"
}

func (c *Commemts) BeforeCreate() (err error) {
	if c.DescriptionId == "" {
		c.DescriptionId = uuid.New().String()
	}
	return
}
