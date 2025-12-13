package main

import (
	"truthly/internals/controller"
	"truthly/internals/repository"
	"truthly/internals/routes"
	"truthly/internals/service"
	"truthly/internals/util"
	"truthly/internals/util/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initialized the logger
	log := logger.InitLogger()
	log.Info("Logger working")

	// Load the enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Error("Error in loading the dotenv file: " + err.Error())
		panic(err.Error())
	}

	// Make the connection with sql
	db := util.InitDb()
	log.Info(db.Name())

	router := gin.Default()

	//User:->  Initialize repo, service, controller
	userRepo := repository.GetUserRepo(log, db)
	userService := service.GetNewUserService(userRepo, log)
	userController := controller.GetNewUserController(userService, log)

	// Register routes
	userRoutes := routes.GetNewUserRoutes(userController)

	userRoutes.RegisterRoutes(router)

	// Post Image
	// repos required
	imageRepo := repository.GetImageRepo(db, log)
	descriptionRepo := repository.GetDescriptionRepository(db, log)
	analyticsRepo := repository.GetAnalyticRepository(db, log)
	commentRepo := repository.GetCommentRepository(db, log)

	s3Uploader, err := service.NewS3Uploader("truthly-images", log)
	if err != nil {
		log.Error(err.Error())
	}

	//services
	postService := service.GetPostService(log, analyticsRepo, commentRepo,
		descriptionRepo, imageRepo, s3Uploader,
	)
	postImageController := controller.GetNewPostImageController(log, postService)

	// register routes
	postImageRoute := routes.GetNewPostImageRoutes(postImageController)

	postImageRoute.RegisterRoutes(router)

	router.Run(":8181")
}
