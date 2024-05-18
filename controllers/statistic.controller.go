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
// @Tags	statistic
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	[]byte
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/statistic/quest/{id} [get]
func (*TStatistic) Quest(context *gin.Context) {
	bearer := utils.GetBearer(context)

	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Csv.GetCsvTableByQuestId(bearer, id)
	utils.CultivateServiceData(context, data, errData)
}
