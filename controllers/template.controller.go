package controllers

import (
	"net/http"

	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TTemplate struct{}

var Template *TTemplate

// Template Get	godoc
// @Summary	Получение шаблона по ID
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.Template
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/template/{id} [get]
func (*TTemplate) Get(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Template.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template GetWithIndicators	godoc
// @Summary	Получение шаблона с вопросами по ID
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.TemplateWithIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/template/with-indicators/{id} [get]
func (*TTemplate) GetWithIndicators(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Template.GetWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template GetAll	godoc
// @Summary	Получение всех шаблонов
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Template
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/template/all [get]
func (*TTemplate) GetAll(context *gin.Context) {
	data, errData := service.Template.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Template Create	godoc
// @Summary	Создание шаблона
// @Tags	template
// @Accept	json
// @Produce	json
// @Param request body models.TemplateCreate true "Body для создания шаблона"
// @Success	200	{object}	models.TemplateWithCountIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/template/create [post]
func (*TTemplate) Create(context *gin.Context) {
	body := &models.TemplateCreate{}

	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	indicators := utils.GetBodyIds(body.Indicators)
	isCondition := utils.CultivateCondition(
		context,
		len(indicators) == 0,
		http.StatusBadRequest,
		"Indicators zero array",
	)
	if isCondition {
		return
	}

	data, count, errData := service.Template.Create(
		body.Name,
		body.Description,
		indicators,
	)
	errService := utils.CultivateServiceError(context, errData)
	if errService != nil {
		return
	}

	returningData := models.TemplateWithCountIndicators{
		IdTemplate:      data.IdTemplate,
		Name:            data.Name,
		Description:     data.Description,
		Available:       data.Available,
		AddedIndicators: count,
	}

	context.JSON(http.StatusOK, returningData)
}

// Template Hide	godoc
// @Summary	Скрытие шаблона по ID (используйте как удаление)
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.Template
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/template/hide/{id} [put]
func (*TTemplate) Hide(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Template.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}
