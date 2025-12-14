package routes

import (
	"truthly/internals/controller"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController *controller.UserController
}

func GetNewUserRoutes(c *controller.UserController) *UserRoutes {
	return &UserRoutes{
		userController: c,
	}
}

func (ur *UserRoutes) RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/v1/users")

	// Create a new user
	userGroup.POST("/", ur.userController.CreateNewUser)
}
