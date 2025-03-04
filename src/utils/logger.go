package utils

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	logFile *os.File
}

var globalLogger *Logger

func NewLogger(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("No se pudo abrir archivo", err)
		return err
	}
	log.SetOutput(file)
	globalLogger = &Logger{logFile: file}
	return nil
}

func (l *Logger) Log(level string, message string) {
	l.logFile.WriteString(fmt.Sprintf("[%s] %s\n", level, message))
}

func GetLogger() *Logger {
	if globalLogger == nil {
		fmt.Println("No se ha iniciado el logger")
	}
	return globalLogger
}
