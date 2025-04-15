package logger

import (
	"log"
)

func Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

func Error(err error) {
	log.Printf("[ERROR] %s", err)
}
