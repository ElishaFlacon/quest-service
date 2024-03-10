package config

import (
	"log"

	"github.com/joho/godotenv"
)

func (init *TInit) InitDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
