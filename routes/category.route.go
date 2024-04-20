package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Category(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/category")
	group.GET("/:id", controllers.Category.Get)
	group.GET("/all", controllers.Category.GetAll)
	group.POST("/create/:id", controllers.Category.Create)
	group.DELETE("/delete/:id", controllers.Category.Delete)
}