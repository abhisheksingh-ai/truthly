package main

import (
	"truthly/internals/app"
	"truthly/internals/util/logger"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.InitLogger()
	log.Info("Logger initialized")

	if err := godotenv.Load(); err != nil {
		log.Error("Error loading env file: " + err.Error())
		panic(err)
	}

	app.Start(log)
}
