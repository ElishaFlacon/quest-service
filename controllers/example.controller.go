package controllers

import (
	"github.com/ElishaFlacon/questionnaire-service/service"
	"github.com/gin-gonic/gin"
)

type TExample struct {
	Value string
}

func (init *TInit) GetExample(context *gin.Context) {
	data, err := service.Init(init.db).GetExample()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (init *TInit) SetExample(context *gin.Context) {
	var body TExample

	errBody := context.BindJSON(&body)
	if errBody != nil {
		context.JSON(500, gin.H{"error": errBody.Error()})
		return
	}

	err := service.Init(init.db).SetExample(body.Value)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": body})
}
