package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	routes := []func(app *gin.Engine){
		Example,
	}

	for index := 0; len(routes) > index; index++ {
		routes[index](app)
	}
}
