package config

import (
	"github.com/ElishaFlacon/quest-service/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUsings(app *gin.Engine) {
	app.Use(middlewares.Cors())
}
