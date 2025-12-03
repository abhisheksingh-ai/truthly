package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId    string    `gorm:"column:UserId;primaryKey"`
	FirstName string    `gorm:"column:FirstName"`
	LastName  string    `gorm:"column:LastName"`
	Age       string    `gorm:"column:Age"`
	Gender    string    `gorm:"column:Gender"`
	Country   string    `gorm:"column:Country"`
	State     string    `gorm:"column:State"`
	City      string    `gorm:"column:City"`
	Address   string    `gorm:"column:Address"`
	Email     string    `gorm:"column:Email"`
	Password  string    `gorm:"column:Password"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

func (User) TableName() string {
	return "User"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserId == "" {
		u.UserId = uuid.New().String()
	}
	return nil
}
