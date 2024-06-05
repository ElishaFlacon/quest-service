package config

import (
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	InitProxies(app)
	InitUsings(app)
}
