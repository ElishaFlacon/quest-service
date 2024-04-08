package routes

import (
	"github.com/ElishaFlacon/quest-service/controllers"
	"github.com/gin-gonic/gin"
)

func Example(app *gin.Engine) {
	example := app.Group("/example")
	example.GET("/all", controllers.Example.GetAll)
	example.POST("/create", controllers.Example.Create)
	example.PUT("/update", controllers.Example.Update)
}
