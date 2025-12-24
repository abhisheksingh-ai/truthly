package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once sync.Once
	log  *slog.Logger
)

func InitLogger() *slog.Logger {
	once.Do(func() {
		// creat a folder and file of app.json
		os.MkdirAll("logs", 0755)

		// create a file and give the excess
		file, err := os.OpenFile("logs/app.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Cannot open log file" + err.Error())
		}

		handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		})

		log = slog.New(handler)
	})
	return log
}
