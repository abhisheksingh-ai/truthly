package routes

import (
	"truthly/internals/controller"
	"truthly/internals/middleware"
	"truthly/internals/util/auth"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController *controller.UserController
	authMiddleware gin.HandlerFunc
}

func GetNewUserRoutes(uc *controller.UserController,
	authToken *auth.AuthToken,
) *UserRoutes {
	return &UserRoutes{
		userController: uc,
		authMiddleware: middleware.AuthMiddleware(authToken),
	}
}

func (ur *UserRoutes) RegisterRoutes(router *gin.RouterGroup) {
	// user group
	userGroup := router.Group("/user")

	userGroup.POST("", ur.authMiddleware, ur.userController.GetUserDetails)
}
