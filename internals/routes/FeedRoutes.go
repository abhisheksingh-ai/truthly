package routes

import (
	"truthly/internals/controller"

	"github.com/gin-gonic/gin"
)

type FeedRoutes struct {
	feedController *controller.FeedController
}

func GetNewFeedRoutes(feedController *controller.FeedController) *FeedRoutes {
	return &FeedRoutes{
		feedController: feedController,
	}
}

func (fr *FeedRoutes) RegisterRoutes(router *gin.Engine) {

	// versioned group
	feedGroup := router.Group("/api/v1/feed")

	// get feed
	feedGroup.GET("/", fr.feedController.GetFeed)
}
