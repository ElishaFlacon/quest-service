package controllers

import (
	"github.com/ElishaFlacon/quest-service/cruds"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/gin-gonic/gin"
)

type TExample struct{}

var Example *TExample

func (*TExample) GetAll(context *gin.Context) {
	data, err := cruds.Example.GetAll()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (*TExample) Create(context *gin.Context) {
	var body models.ExampleBody

	errBody := context.BindJSON(&body)

	if errBody != nil {
		context.JSON(500, gin.H{"error": errBody.Error()})
		return
	}

	data, err := cruds.Example.Create(body.Data)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}

func (*TExample) Update(context *gin.Context) {
	var body models.Example

	errBody := context.BindJSON(&body)

	if errBody != nil {
		context.JSON(500, gin.H{"error": errBody.Error()})
		return
	}

	data, err := cruds.Example.Update(body.Id, body.Value)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}
