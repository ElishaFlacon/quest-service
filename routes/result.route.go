package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Result(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/result")
	group.GET("/by-user/:id", controllers.Result.GetByUserId)
	group.GET("/by-users", controllers.Result.GetByUsersId)
	group.GET("/by-quest/:id", controllers.Result.GetByQuestId)
	group.GET("/by-quest/:id_quest/by_team/:id_team", controllers.Result.GetByQuestIdAndTeamId)
	group.POST("/create", controllers.Result.Create)
}
