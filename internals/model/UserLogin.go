package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLogin struct {
	Id          string    `gorm:"column:Id"`
	UserId      string    `gorm:"column:UserId"`
	Name        string    `gorm:"column:Name"`
	Email       string    `gorm:"column:Email"`
	LastLogInAt string    `gorm:"column:LastLogInAt"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	UpdatedAt   time.Time `gorm:"column:UpdatedAt"`
}

func (UserLogin) TableName() string {
	return "UserLogIn"
}

func (u *UserLogin) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return nil
}
