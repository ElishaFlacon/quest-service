package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Category(app *gin.Engine) {
	group := app.Group("/category")
	group.GET("/:id", controllers.Category.Get)

	// TODO доделать все роуты
}
