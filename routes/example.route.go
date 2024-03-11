package routes

import (
	"github.com/ElishaFlacon/questionnaire-service/controllers"
	"github.com/gin-gonic/gin"
)

func Example(app *gin.Engine) {
	app.GET("/example", controllers.Example.Get)
	app.POST("/example", controllers.Example.Set)
}
