package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Template(app *gin.Engine) {
	template := app.Group("/template")
	template.GET("/:id", controllers.Template.Get)
	template.GET("/all", controllers.Template.GetAll)
	template.POST("/create", controllers.Template.Create)
}
