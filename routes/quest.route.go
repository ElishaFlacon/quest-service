package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Quest(app *gin.Engine) {
	group := app.Group("/quest")
	group.GET("/:id", controllers.Quest.Get)

	// TODO доделать все роуты
}
