package routes

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/gin-gonic/gin"
)

type TInit struct {
	app *gin.Engine
	db  *database.TInit
}

func Init(app *gin.Engine, db *database.TInit) *TInit {
	return &TInit{app: app, db: db}
}

func (init *TInit) Start() {
	routes := []func(){
		init.ExampleRoutes,
	}

	for index := 0; len(routes) > index; index++ {
		routes[index]()
	}
}
