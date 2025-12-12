package controller

import (
	"log/slog"
	"truthly/internals/service"

	"github.com/gin-gonic/gin"
)

type PostImageController struct {
	logger      *slog.Logger
	postService service.PostService
}

// constructor
func GetNewPostImageController(logger *slog.Logger, postService service.PostService) *PostImageController {
	return &PostImageController{
		logger:      logger,
		postService: postService,
	}
}

func (h *PostImageController) PostImage(c *gin.Context) {

}
