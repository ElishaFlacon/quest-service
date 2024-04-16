package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Docs(app *gin.Engine) {
	group := app.Group("/docs")
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
