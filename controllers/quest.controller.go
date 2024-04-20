package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/gin-gonic/gin"
)

// Таблица
// Деталка сам опрос + участники
// Выдача вопросов (частично сделано в индикаторах)
// Создание
// удаление\скрытие (в самом конце)

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

// Quest Get	godoc
// @Summary	Пример get quest
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/:id [get]
func (*TQuest) Get(context *gin.Context) {
	// TODO
}
