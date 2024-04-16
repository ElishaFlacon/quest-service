package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	routes := []func(app *gin.Engine){
		Category,
		Docs,
		Indicator,
		Quest,
		Result,
		Template,
	}

	for index := range routes {
		routes[index](app)
	}
}
