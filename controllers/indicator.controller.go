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
// @Param	id path int true "ID вопроса"
// @Success	200	{object}	models.IndicatorWithCategoryName
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/indicator/{id} [get]
func (*TIndicator) Get(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Indicator.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetByTemplateId	godoc
// @Summary	Получение вопросов по ID шаблона
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{array}	models.IndicatorWithCategoryName
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/indicator/by-template/{id} [get]
func (*TIndicator) GetByTemplateId(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Indicator.GetByTemplateId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetByQuestId	godoc
// @Summary	Получение вопросов по ID опроса
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{array}	models.IndicatorWithCategoryName
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/indicator/by-quest/{id} [get]
func (*TIndicator) GetByQuestId(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Indicator.GetByQuestId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetAll	godoc
// @Summary	Получение всех вопросов
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.IndicatorWithCategoryName
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
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
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/indicator/create [post]
func (*TIndicator) Create(context *gin.Context) {
	body := &models.IndicatorCreateRequest{}

	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	data, errData := service.Indicator.Create(
		body.IdCategory,
		body.Name,
		body.Description,
		body.FromRole,
		body.ToRole,
	)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Hide	godoc
// @Summary	Скрытие вопроса по ID
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Param	id path int true "ID вопроса"
// @Success	200	{object}	models.Indicator
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/indicator/hide/{id} [put]
func (*TIndicator) Hide(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Indicator.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Delete	godoc
// @Summary	Удаление вопроса по ID (нельзя удалить, если вопрос используется)
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Param	id path int true "ID вопроса"
// @Success	200	{object}	models.Indicator
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/indicator/delete/{id} [delete]
func (*TIndicator) Delete(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Indicator.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
