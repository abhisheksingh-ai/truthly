package service

import (
	"context"
	"log/slog"
	"truthly/internals/dto"
	"truthly/internals/repository"
)

// Interface
type UserService interface {
	CreateNewUser(ctx context.Context, userReq *dto.UserRequestDto) (*dto.ResponseDto[any], error)
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

func (s *userService) CreateNewUser(ctx context.Context, userReq *dto.UserRequestDto) (*dto.ResponseDto[any], error) {
	//1. dto -> model
	user := dto.ToModel(userReq)

	//2. service will call to repo
	savedUser, err := s.userRepo.CreatNewUser(ctx, user)
	if err != nil {
		s.logger.Error("Error while creating a new user, Error: " + err.Error())
		return nil, err
	}

	s.logger.Info("New user created", "userId", savedUser.UserId)

	//3. model -> dto
	return &dto.ResponseDto[any]{
		Status:  "success",
		Message: "User created",
		ResultObj: map[string]interface{}{
			"userId": savedUser.UserId,
		},
	}, nil
}
