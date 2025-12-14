package routes

import (
	"log/slog"
	"truthly/internals/controller"
	"truthly/internals/repository"
	"truthly/internals/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAll(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {

	registerUser(router, db, logger)
	registerPost(router, db, logger)
	registerFeed(router, db, logger)
}

// user
func registerUser(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {

	userRepo := repository.GetUserRepo(logger, db)
	userService := service.GetNewUserService(userRepo, logger)
	userController := controller.GetNewUserController(userService, logger)

	GetNewUserRoutes(userController).RegisterRoutes(router)
}

// post
func registerPost(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {

	imageRepo := repository.GetImageRepo(db, logger)
	descriptionRepo := repository.GetDescriptionRepository(db, logger)
	analyticsRepo := repository.GetAnalyticRepository(db, logger)
	commentRepo := repository.GetCommentRepository(db, logger)

	s3Uploader, err := service.NewS3Uploader("truthly-images", logger)
	if err != nil {
		logger.Error(err.Error())
	}

	postService := service.GetPostService(logger, analyticsRepo, commentRepo,
		descriptionRepo, imageRepo, s3Uploader,
	)
	postImageController := controller.GetNewPostImageController(logger, postService)

	GetNewPostImageRoutes(postImageController).RegisterRoutes(router)

}

// feed
func registerFeed(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {

	feedRepo := repository.GetNewFeedRepository(db, logger)
	commentRepo := repository.GetCommentRepository(db, logger)
	feedService := service.GetNewFeedService(feedRepo, commentRepo, logger)
	feedController := controller.GetNewFeedController(logger, feedService)

	GetNewFeedRoutes(feedController).RegisterRoutes(router)
}
