package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Indicator(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/indicator")
	group.GET("/:id", controllers.Indicator.Get)
	group.GET("/by-template/:id", controllers.Indicator.GetByTemplateId)
	group.GET("/by-quest/:id", controllers.Indicator.GetByQuestId)
	group.GET("/all", controllers.Indicator.GetAll)
	group.POST("/create", controllers.Indicator.Create)
	group.PUT("/hide/:id", controllers.Indicator.Hide)
	group.DELETE("/delete/:id", controllers.Indicator.Delete)
}
