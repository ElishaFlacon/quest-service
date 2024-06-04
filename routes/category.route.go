package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Category(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/category")
	group.GET("/:id", controllers.Category.Get)
	group.GET("/all", controllers.Category.GetAll)
	group.POST("/create", controllers.Category.Create)
	group.PUT("/hide/:id", controllers.Category.Hide)
}
