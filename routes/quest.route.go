package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Quest(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/quest")
	group.GET("/:id", controllers.Quest.Get)
	group.GET("/with-indicators/:id", controllers.Quest.GetWithIndicators)
	group.GET("/with-users/:id", controllers.Quest.GetWithUsers)
	group.GET("/with-users-and-indicators/:id", controllers.Quest.GetWithUsersAndIndicators)
	group.GET("/all", controllers.Quest.GetAll)
	group.POST("/create", controllers.Quest.Create)
	group.PUT("/hide/:id", controllers.Quest.Hide)
	group.DELETE("/delete/:id", controllers.Quest.Delete)
}
