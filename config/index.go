package config

import (
	"github.com/gin-gonic/gin"
)

type TInit struct {
	app *gin.Engine
}

func Init(app *gin.Engine) *TInit {
	return &TInit{app: app}
}

func (init *TInit) Start() {
	init.InitDotenv()
	init.InitProxies()
}
