package utils

import (
	"os"
)

func GetAppMode() string {
	mode := os.Getenv("APP_MODE")

	if mode == "PRODUCTION" {
		return "release"
	}

	return "debug"
}
