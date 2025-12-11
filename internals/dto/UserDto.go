package dto

import (
	"truthly/internals/model"
)

type UserRequestDto struct {
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	UserName     string `json:"userName,omitempty"`
	Age          int    `json:"age,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Country      string `json:"country,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Address      string `json:"address,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	MobileNumber string `json:"mobileNumber,omitempty"`
}

type UserResponseDto struct {
	Message string `json:"message,omitempty"`
	UserId  string `json:"userId"`
}

// DTO → Model

func ToModel(u *UserRequestDto) *model.User {
	return &model.User{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		UserName:     u.UserName,
		Age:          u.Age,
		Gender:       u.Gender,
		Country:      u.Country,
		State:        u.State,
		City:         u.City,
		Address:      u.Address,
		Email:        u.Email,
		Password:     u.Password, // later hash this in service
		MobileNumber: u.MobileNumber,
	}
}

// Model to dto I will do manually
