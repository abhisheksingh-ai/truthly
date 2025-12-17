package service

import (
	"context"
	"log/slog"
	"time"
	"truthly/internals/dto"
	"truthly/internals/model"
	"truthly/internals/repository"
)

type AuthService interface {
	UserSignup(ctx context.Context, user *dto.UserRequestDto) (*dto.ResponseDto[*dto.LogInRes], error)
	VerifyMail(ctx context.Context, loginReq *dto.LoginReq) (*dto.ResponseDto[*dto.LogInRes], error)
	AddSession(ctx context.Context, sessionId string, userId string, token string) (*dto.ResponseDto[*dto.LogInRes], error)
}

type authService struct {
	logger          *slog.Logger
	userLoginRepo   repository.UserLoginRepository
	userSessionRepo repository.UserSessionRepository
	userRepo        repository.UserRepository
}

func GetNewAuthService(logger *slog.Logger, userLoginRepo repository.UserLoginRepository, userSessionRepo repository.UserSessionRepository, userRepo repository.UserRepository) AuthService {
	return &authService{
		logger:          logger,
		userLoginRepo:   userLoginRepo,
		userSessionRepo: userSessionRepo,
		userRepo:        userRepo,
	}
}

func (s *authService) UserSignup(ctx context.Context, userReq *dto.UserRequestDto) (*dto.ResponseDto[*dto.LogInRes], error) {
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
	return &dto.ResponseDto[*dto.LogInRes]{
		Status:  "success",
		Message: "User created",
		ResultObj: &dto.LogInRes{
			UserId: savedUser.UserId,
		},
	}, nil
}

// verify mail before log in and get the user id too
func (s *authService) VerifyMail(ctx context.Context, loginReq *dto.LoginReq) (*dto.ResponseDto[*dto.LogInRes], error) {
	// mail
	email := loginReq.Email

	// call repo and verfiy
	res, err := s.userRepo.VerifyMail(ctx, email)
	if err != nil {
		s.logger.Error(err.Error())
		return &dto.ResponseDto[*dto.LogInRes]{
			Status:    "Error",
			Error:     err.Error(),
			ResultObj: nil,
		}, err
	}

	// if you get the user
	return &dto.ResponseDto[*dto.LogInRes]{
		Status:  "Success",
		Message: "User exists",
		ResultObj: &dto.LogInRes{
			UserId:   res.UserId,
			UserName: res.UserName,
		},
	}, nil
}

// Add session
func (s *authService) AddSession(ctx context.Context, sessionId string, userId string, token string) (*dto.ResponseDto[*dto.LogInRes], error) {
	// data ---> model
	userSession := model.UserSession{
		UserId:    userId,
		SessionId: sessionId,
		Status:    "ACTIVE",
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
	// repo calling
	err := s.userSessionRepo.CreateNewSession(ctx, &userSession)
	if err != nil {
		return &dto.ResponseDto[*dto.LogInRes]{
			Error: err.Error(),
		}, nil
	}

	// return
	return &dto.ResponseDto[*dto.LogInRes]{
		Status: "Token created",
		ResultObj: &dto.LogInRes{
			Token: token,
		},
	}, nil
}
