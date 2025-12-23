package routes

import (
	"log/slog"
	"truthly/internals/controller"
	"truthly/internals/realtime"
	"truthly/internals/repository"
	"truthly/internals/service"
	"truthly/internals/util/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAll(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger, hub *realtime.Hub) {
	registerPost(router, db, logger)
	registerFeed(router, db, logger)
	registerAuth(router, db, logger)
	registerInteraction(router, db, logger)
	registerWebsocket(router, hub, db, logger)
}

// post
func registerPost(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {

	imageRepo := repository.GetImageRepo(db, logger)
	descriptionRepo := repository.GetDescriptionRepository(db, logger)
	analyticsRepo := repository.GetAnalyticRepository(db, logger)
	commentRepo := repository.GetCommentRepository(db, logger)
	userSessionRepo := repository.GetNewUserSessionRepo(logger, db)

	s3Uploader, err := service.NewS3Uploader("truthly-images", logger)
	if err != nil {
		logger.Error(err.Error())
	}

	postService := service.GetPostService(logger, analyticsRepo, commentRepo,
		descriptionRepo, imageRepo, s3Uploader,
	)
	postImageController := controller.GetNewPostImageController(logger, postService)

	authToken := auth.GetNewAuthToken(logger, userSessionRepo)

	GetNewPostImageRoutes(postImageController, authToken).RegisterRoutes(router)

}

// feed
func registerFeed(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {

	feedRepo := repository.GetNewFeedRepository(db, logger)
	commentRepo := repository.GetCommentRepository(db, logger)
	feedService := service.GetNewFeedService(feedRepo, commentRepo, logger)
	feedController := controller.GetNewFeedController(logger, feedService)
	userSessionRepo := repository.GetNewUserSessionRepo(logger, db)

	authToken := auth.GetNewAuthToken(logger, userSessionRepo)

	GetNewFeedRoutes(feedController, authToken).RegisterRoutes(router)
}

// auth
func registerAuth(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {
	// repo's required
	userLoginRepo := repository.GetNewUserLoginRepo(logger, db)
	userSessionRepo := repository.GetNewUserSessionRepo(logger, db)
	userRepo := repository.GetUserRepo(logger, db)

	// auth service
	authService := service.GetNewAuthService(logger, userLoginRepo, userSessionRepo, userRepo)

	// utils
	authUtil := auth.GetNewAuthToken(logger, userSessionRepo)

	// auth controller
	authController := controller.GetNewAuthController(logger, authService, authUtil)

	// routes
	GetNewAuthRoutes(authController).RegisterRoutes(router)
}

// interaction
func registerInteraction(router *gin.RouterGroup, db *gorm.DB, logger *slog.Logger) {
	// repo required
	interactionRepo := repository.GetNewInteractionRepository(db, logger)
	analyticRepo := repository.GetAnalyticRepository(db, logger)
	userSessionRepo := repository.GetNewUserSessionRepo(logger, db)

	// service required
	interactionService := service.GetNewInteractionService(logger, interactionRepo, analyticRepo)

	// auth
	authToken := auth.GetNewAuthToken(logger, userSessionRepo)

	// controller
	interactionController := controller.GetNewInteractionController(logger, interactionService)

	// routes
	GetNewInteractionRoutes(interactionController, authToken).RegisterRoutes(router)
}

// websocket
// websocket
func registerWebsocket(
	router *gin.RouterGroup,
	hub *realtime.Hub,
	db *gorm.DB,
	logger *slog.Logger,
) {

	// repo
	userSessionRepo := repository.GetNewUserSessionRepo(logger, db)

	// auth token
	authToken := auth.GetNewAuthToken(logger, userSessionRepo)

	// route group
	wsGroup := router.Group("/ws")

	wsGroup.GET("/", func(c *gin.Context) {
		controller.ServeWS(
			hub,
			c.Writer,
			c.Request,
			authToken,
		)
	})
}
