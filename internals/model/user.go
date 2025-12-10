package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId string `gorm:"column:UserId;primaryKey"`

	UserName string `gorm:"column:UserName;unique;size:20;not null"`

	FirstName string `gorm:"column:FirstName;size:50;not null"`
	LastName  string `gorm:"column:LastName;size:50;not null"`

	Age int `gorm:"column:Age;check:age_check,age >= 0 AND age <= 120"`

	Gender string `gorm:"column:Gender;size:10;check:gender_check,gender IN ('male','female','other')"`

	Country string `gorm:"column:Country;size:56;not null"`
	State   string `gorm:"column:State;size:56;not null"`
	City    string `gorm:"column:City;size:56;not null"`
	Address string `gorm:"column:Address;size:255"`

	Email        string `gorm:"column:Email;unique;size:100;not null"`
	Password     string `gorm:"column:Password;size:255;not null"` // hash stored
	MobileNumber string `gorm:"column:MobileNumber;unique;size:15;not null"`

	CreatedAt time.Time `gorm:"column:CreatedAt;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;autoUpdateTime"`
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
