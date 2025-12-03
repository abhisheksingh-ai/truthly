package main

import "truthly/internals/util/logger"

func main() {
	log := logger.InitLogger()
	log.Info("Logger working")
}
