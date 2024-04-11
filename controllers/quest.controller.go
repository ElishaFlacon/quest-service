package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/gin-gonic/gin"
)

type TQuest struct{}

var Quest *TQuest

func (*TQuest) GetAll(context *gin.Context) {
	data, err := service.Quest.GetAll()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": data})
}
