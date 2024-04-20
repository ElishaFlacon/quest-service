package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Result(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/result")
	group.GET("/user/:id", controllers.Result.GetByUserId)
	group.GET("/users", controllers.Result.GetByUsersId)
	group.POST("/create", controllers.Result.Create)
}
