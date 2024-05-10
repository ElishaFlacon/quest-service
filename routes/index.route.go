package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	group := app.Group("/quest-service")

	routes := []func(app *gin.RouterGroup){
		Category,
		Docs,
		Indicator,
		Quest,
		Result,
		Template,
	}

	for index := range routes {
		routes[index](group)
	}
}
