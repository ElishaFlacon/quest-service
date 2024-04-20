package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Indicator(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/indicator")
	group.GET("/:id", controllers.Indicator.Get)
	group.GET("/all", controllers.Indicator.GetAll)
	group.GET("/quest/:id", controllers.Indicator.GetByQuestId)
	group.POST("/create", controllers.Indicator.Create)
}
