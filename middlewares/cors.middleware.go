package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)
		context.Writer.Header().Set(
			"Access-Control-Allow-Credentials",
			"true",
		)
		context.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			`
				Content-Type, 
				Content-Length, 
				Accept-Encoding, 
				X-CSRF-Token, 
				Authorization, 
				accept, 
				origin, 
				Cache-Control, 
				X-Requested-With
			`,
		)
		context.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}

		context.Next()
	}
}
