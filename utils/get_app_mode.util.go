package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GetAppMode() string {
	mode := os.Getenv("APP_MODE")

	if mode == "PRODUCTION" {
		return gin.ReleaseMode
	}

	return gin.DebugMode
}
