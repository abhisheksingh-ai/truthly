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

	router := gin.Default()

	// versioned API
	api := router.Group("/api/v1")

	// register all routes
	routes.RegisterAll(api, db, logger)

	router.Run(":8181")
}
