package middlewares

import "github.com/gin-gonic/gin"

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// тут будет отправка запроса на проверку токена
		// не будет :3
		c.Next()
	}
}
