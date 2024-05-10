package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Template(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/template")
	group.GET("/:id", controllers.Template.Get)
	group.GET("/with-indicators/:id", controllers.Template.GetWithIndicators)
	group.GET("/all", controllers.Template.GetAll)
	group.POST("/create", controllers.Template.Create)
	group.PUT("/hide/:id", controllers.Template.Hide)
	group.DELETE("/delete/:id", controllers.Template.Delete)
}
