package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	err := godotenv.Load()

	if err == nil {
		return
	}

	log.Fatal("Error loading .env file: ", err)
}
