package utils

import (
	"os"
)

func GetUserServiceUrl() string {
	return os.Getenv("USER_SERVICE_URL")
}
