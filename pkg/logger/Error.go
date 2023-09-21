package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func Error(message string) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	os.Stderr = file

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Error(message)
}
