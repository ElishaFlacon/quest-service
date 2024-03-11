package routes

import (
	"github.com/ElishaFlacon/questionnaire-service/controllers"
	"github.com/gin-gonic/gin"
)

func Example(app *gin.Engine) {
	app.GET("/example", controllers.GetExample)
	app.POST("/example", controllers.SetExample)
}
