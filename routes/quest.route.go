package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Quest(app *gin.Engine) {
	quest := app.Group("/quest")
	quest.GET("/all", controllers.Quest.GetAll)
}
