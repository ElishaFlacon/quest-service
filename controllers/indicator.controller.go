package controllers

import (
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TIndicator struct{}

var Indicator *TIndicator

// Indicator Get	godoc
// @Summary	Получение вопроса по ID
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/{id} [get]
func (*TIndicator) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetByTemplateId	godoc
// @Summary	Получение вопросов по ID шаблона
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/by-template/{id} [get]
func (*TIndicator) GetByTemplateId(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.GetByTemplateId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetByQuestId	godoc
// @Summary	Получение вопросов по ID опроса
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/by-quest/{id} [get]
func (*TIndicator) GetByQuestId(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.GetByQuestId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetAll	godoc
// @Summary	Получение всех вопросов
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/all [get]
func (*TIndicator) GetAll(context *gin.Context) {
	data, errData := service.Indicator.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Create	godoc
// @Summary	Создание вопроса
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Param request body models.IndicatorCreateRequest true "Body для создания вопроса"
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/create [post]
func (*TIndicator) Create(context *gin.Context) {
	body := models.IndicatorCreateRequest{}
	utils.CultivateBody(context, body)
	data, errData := service.Indicator.Create(
		body.Name,
		body.Description,
		body.Role,
		body.Visible,
		body.IdCategory,
	)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Hide	godoc
// @Summary	Скрытие вопроса по ID
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/hide/{id} [put]
func (*TIndicator) Hide(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Delete	godoc
// @Summary	Удаление вопроса по ID
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/indicator/delete/{id} [delete]
func (*TIndicator) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
