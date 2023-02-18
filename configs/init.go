package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var File *os.File

func Init() {
	file, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	File = file

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
