package utils

import (
	"os"
)

func GetProfileServiceUrl() string {
	return os.Getenv("PROFILE_SERVICE_URL")
}
