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
	registerPost(router, db, logger)
	registerFeed(router, db, logger)
	registerAuth(router, db, logger)
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

// auth
func registerAuth(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {
	// repo's required
	userLoginRepo := repository.GetNewUserLoginRepo(logger, db)
	userSessionRepo := repository.GetNewUserSessionRepo(logger, db)
	userRepo := repository.GetUserRepo(logger, db)

	// auth service
	authService := service.GetNewAuthService(logger, userLoginRepo, userSessionRepo, userRepo)

	// auth controller
	authController := controller.GetNewAuthController(logger, authService)

	// routes
	GetNewAuthRoutes(authController).RegisterRoutes(router)
}
