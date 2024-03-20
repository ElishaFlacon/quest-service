package config

import "github.com/gin-gonic/gin"

func InitProxies(app *gin.Engine) {
	app.ForwardedByClientIP = true
	app.SetTrustedProxies([]string{
		"localhost",
		"127.0.0.1",
	})
}
