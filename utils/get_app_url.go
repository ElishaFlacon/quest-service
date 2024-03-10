package utils

import (
	"os"
)

func GetAppUrl() string {
	return os.Getenv("APP_URL")
}
