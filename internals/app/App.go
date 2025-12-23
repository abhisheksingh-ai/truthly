package app

import (
	"log/slog"
	"truthly/internals/realtime"
	"truthly/internals/routes"
	"truthly/internals/util"

	"github.com/gin-gonic/gin"
)

func Start(logger *slog.Logger) {
	db := util.InitDb()
	logger.Info("DB Connected")

	sqlDb, err := db.DB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer sqlDb.Close() // Close db connection when the app is shut down

	router := gin.Default()

	// health route if my app is runnig
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "app is running",
		})
	})

	// hub
	hub := realtime.NewHub()
	go hub.Run()

	// versioned API
	api := router.Group("/api/v1")

	// register all routes
	routes.RegisterAll(api, db, logger, hub)

	router.Run(":8181")
}
