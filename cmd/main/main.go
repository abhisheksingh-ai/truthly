package main

import (
	"truthly/internals/util"
	"truthly/internals/util/logger"

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
}
