package config

import (
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	InitDotenv()
	InitProxies(app)
	InitUsings(app)
}
