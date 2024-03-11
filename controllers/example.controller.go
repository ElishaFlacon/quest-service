package controllers

import (
	"github.com/ElishaFlacon/questionnaire-service/service"
	"github.com/gin-gonic/gin"
)

type TExampleBody struct {
	Value string
}

func GetExample(context *gin.Context) {
	data, err := service.GetExample()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func SetExample(context *gin.Context) {
	var body TExampleBody

	errBody := context.BindJSON(&body)
	if errBody != nil {
		context.JSON(500, gin.H{"error": errBody.Error()})
		return
	}

	err := service.SetExample(body.Value)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": body})
}
