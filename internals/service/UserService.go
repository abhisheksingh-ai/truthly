package service

import (
	"context"
	"log/slog"
	"truthly/internals/dto"
	"truthly/internals/repository"
)

type UserService interface {
	GetUserById(ctx context.Context, userId string) (*dto.UserDetailsForHome, error)
}

type userService struct {
	userRepo repository.UserRepository
	logger   *slog.Logger
}

func GetNewUserService(logger *slog.Logger, ur repository.UserRepository) UserService {
	return &userService{
		logger:   logger,
		userRepo: ur,
	}
}

func (us *userService) GetUserById(ctx context.Context, userId string) (*dto.UserDetailsForHome, error) {

	user, err := us.userRepo.GetUserById(ctx, userId)

	if err != nil {
		us.logger.Error("Failed to fetch the user", "userId", userId)
		return nil, err
	}

	resp := &dto.UserDetailsForHome{
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,

		City:    user.City,
		State:   user.State,
		Country: user.Country,
	}

	return resp, nil
}
