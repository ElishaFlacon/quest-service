package controllers

import "github.com/gin-gonic/gin"

// Таблица
// Деталка сам опрос + участники
// Выдача вопросов (частично сделано в индикаторах)
// Создание
// удаление\скрытие (в самом конце)

type TQuest struct{}

var Quest *TQuest

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
