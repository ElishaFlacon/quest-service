package utils

import "github.com/gin-gonic/gin"

func GetBearer(context *gin.Context) string {
	bearer := context.GetHeader("Authorization")
	return bearer
}
