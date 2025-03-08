package utils

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(logPath string) {
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("No se pudo abrir el archivo del log")
	}
	fileHandler := slog.NewTextHandler(file, nil)
	logger := slog.New(fileHandler)
	slog.SetDefault(logger)
}
