package routes

import (
	"github.com/ElishaFlacon/questionnaire-service/controllers"
	"github.com/gin-gonic/gin"
)

func Template(app *gin.Engine) {
	template := app.Group("/template")
	template.POST("/create", controllers.Template.Create)
}
