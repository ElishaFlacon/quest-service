package routes

import (
	"github.com/ElishaFlacon/questionnaire-service/controllers"
	"github.com/gin-gonic/gin"
)

func Indicator(app *gin.Engine) {
	indicator := app.Group("/indicator")
	indicator.GET("/{id}", controllers.Indicator.Get)
	indicator.GET("/all", controllers.Indicator.GetAll)
	indicator.GET("/all/for-quest", controllers.Indicator.GetAllForQuest)
	indicator.POST("/create", controllers.Example.Create)
}
