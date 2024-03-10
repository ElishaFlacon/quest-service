package routes

import (
	"github.com/ElishaFlacon/questionnaire-service/controllers"
)

func (init *TInit) ExampleRoutes() {
	init.app.GET("/example", controllers.Init(init.db).GetExample)
	init.app.POST("/example", controllers.Init(init.db).SetExample)
}
