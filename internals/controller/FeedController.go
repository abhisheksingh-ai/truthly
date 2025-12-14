package controller

import (
	"log/slog"
	"strconv"
	"truthly/internals/service"

	"github.com/gin-gonic/gin"
)

type FeedController struct {
	logger      *slog.Logger
	feedService service.FeedService
}

func GetNewFeedController(l *slog.Logger, fs service.FeedService) *FeedController {
	return &FeedController{
		logger:      l,
		feedService: fs,
	}
}

func (c *FeedController) GetFeed(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	cursor := ctx.Query("cursor")

	resp, err := c.feedService.GetFeed(ctx, limit, cursor)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error,
		})
	}

	ctx.JSON(200, resp)
}
