package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Indicator(app *gin.Engine) {
	indicator := app.Group("/indicator")
	indicator.GET("/:id", controllers.Indicator.Get)
	indicator.GET("/all", controllers.Indicator.GetAll)
	indicator.GET("/quest/:id", controllers.Indicator.GetByQuestId)
	indicator.POST("/create", controllers.Indicator.Create)
}
