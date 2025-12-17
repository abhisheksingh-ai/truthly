package routes

import (
	"truthly/internals/controller"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController *controller.AuthController
}

func GetNewAuthRoutes(authController *controller.AuthController) *AuthRoutes {
	return &AuthRoutes{
		authController: authController,
	}
}

func (ar *AuthRoutes) RegisterRoutes(router *gin.RouterGroup) {
	// grouping auth
	authGroup := router.Group("/auth")

	// signup
	authGroup.POST("/signup", ar.authController.UserSignup)

	// login
	authGroup.POST("/login", ar.authController.UserLogin)
}
