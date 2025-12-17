package controller

import (
	"log/slog"
	"truthly/internals/dto"
	"truthly/internals/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger      *slog.Logger
	authService service.AuthService
}

func GetNewAuthController(logger *slog.Logger, authService service.AuthService) *AuthController {
	return &AuthController{
		logger:      logger,
		authService: authService,
	}
}

func (c *AuthController) UserSignup(ctx *gin.Context) {
	// variable to taking data from request
	var user dto.UserRequestDto

	// binding the data
	if err := ctx.ShouldBindJSON(&user); err != nil {
		c.logger.Error(err.Error())
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// call the service to signup the user
	resp, err := c.authService.UserSignup(ctx, &user)
	if err != nil {
		c.logger.Error(err.Error())
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, resp)
}
