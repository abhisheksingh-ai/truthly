package service

import (
	"context"
	"log/slog"
	"truthly/internals/dto"
	"truthly/internals/repository"
)

// Interface
type UserService interface {
	CreateNewUser(ctx context.Context, user *dto.UserRequestDto) (*dto.UserResponseDto, error)
}

// struct
type userService struct {
	userRepo repository.UserRepository
	logger   *slog.Logger
}

// constructor
func GetNewUserService(r repository.UserRepository, l *slog.Logger) UserService {
	return &userService{
		userRepo: r,
		logger:   l,
	}
}

func (s *userService) CreateNewUser(ctx context.Context, userDto *dto.UserRequestDto) (*dto.UserResponseDto, error) {
	//1. dto -> model
	user := dto.ToModel(userDto)

	//2. service will call to repo
	savedUser, err := s.userRepo.CreatNewUser(ctx, user)
	if err != nil {
		s.logger.Error("Error while creating a new user, Error: " + err.Error())
		return nil, err
	}

	//3. model -> dto
	return &dto.UserResponseDto{
		Message: "User created successfull",
		UserId:  savedUser.UserId,
	}, nil
}
