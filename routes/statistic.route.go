package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Statistic(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/statistic")
	group.GET("/quest/:id", controllers.Statistic.Quest)
}
