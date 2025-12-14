package routes

import (
	"truthly/internals/controller"

	"github.com/gin-gonic/gin"
)

type PostImageRoutes struct {
	postImageController *controller.PostImageController
}

func GetNewPostImageRoutes(postImageController *controller.PostImageController) *PostImageRoutes {
	return &PostImageRoutes{
		postImageController: postImageController,
	}
}

func (p *PostImageRoutes) RegisterRoutes(router *gin.RouterGroup) {

	// group
	postImageGroup := router.Group("/posts")

	// create a new post
	postImageGroup.POST("/", p.postImageController.PostImage)
}
