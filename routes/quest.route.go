package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Quest(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/quest")
	group.GET("/:id", controllers.Quest.Get)
	group.GET("/all", controllers.Quest.GetAll)

	// TODO доделать все роуты
}
