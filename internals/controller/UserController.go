package controller

import (
	"log/slog"
	"net/http"
	"truthly/internals/dto"
	"truthly/internals/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
	logger      *slog.Logger
}

func GetNewUserController(s service.UserService, l *slog.Logger) *UserController {
	return &UserController{
		userService: s,
		logger:      l,
	}
}

func (h *UserController) CreateNewUser(ctx *gin.Context) {

	// variable to taking data from request
	var user dto.UserRequestDto

	// binding the data
	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, &dto.UserResponseDto{
			Message: err.Error(),
		})
		return
	}

	// call the service
	response, err := h.userService.CreateNewUser(ctx, &user)
	if err != nil {
		h.logger.Error(err.Error())
		ctx.JSON(500, &dto.UserResponseDto{
			Message: err.Error(),
		})
		return
	}

	h.logger.Info("New user created", "userId", response.UserId)
	ctx.JSON(200, &dto.UserResponseDto{
		UserId: response.UserId,
	})
}
