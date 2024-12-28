package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvHandler() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
