package controller

import (
	"truthly/internals/dto"
	"truthly/internals/service"

	"github.com/gin-gonic/gin"
)

// http:IP:PORT//api/v1/user/
type UserController struct {
	userService service.UserService
}

func GetNewUserController(us service.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

// api/v1/user/
func (uc *UserController) GetUserDetails(ctx *gin.Context) {

	userId := ctx.GetString("userId")

	userDetails, err := uc.userService.GetUserById(ctx, userId)
	if err != nil {
		ctx.JSON(500, dto.ResponseDto[any]{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(200, dto.ResponseDto[any]{
		Status:    "Success",
		ResultObj: userDetails,
	})
}
