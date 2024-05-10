package middlewares

import "github.com/gin-gonic/gin"

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO тут будет отправка запроса на проверку токена
		c.Next()
	}
}
