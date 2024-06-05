package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Docs(baseGroup *gin.RouterGroup) {
	group := baseGroup.Group("/docs")

	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
