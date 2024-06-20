package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TStatistic struct{}

var Statistic *TStatistic

// Statistic Quest	godoc
// @Summary	Получение статистики опроса по ID опроса
// @Description Ответ содержит файл CSV с колонками: Название опроса, Название вопроса, Название команды, От кого, Роль (от кого), Кому, Роль (кому), Результат ответа
// @Tags	statistic
// @Accept	json
// @Produce	text/csv
// @Param	id path int true "ID опроса"
// @Success	200	{file}	[]byte
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/statistic/quest/{id} [get]
func (*TStatistic) Quest(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Statistic.GetQuestStatistic(id)
	utils.CultivateCsvData(context, data, errData)
}
