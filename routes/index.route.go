package routes

import (
	"github.com/gin-gonic/gin"
)

// При добавлении нового роута, не забываем добавить его сюда

func Init(app *gin.Engine) {
	group := app.Group("/quest-service")

	routes := []func(app *gin.RouterGroup){
		Category,
		Docs,
		Indicator,
		Quest,
		Result,
		Template,
		Statistic,
	}

	for index := range routes {
		routes[index](group)
	}
}
