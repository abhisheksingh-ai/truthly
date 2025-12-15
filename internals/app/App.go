package app

import (
	"log/slog"
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

	// versioned API
	api := router.Group("/api/v1")

	// register all routes
	routes.RegisterAll(api, db, logger)

	router.Run(":8181")
}
