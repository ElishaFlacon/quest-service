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
// @Summary	Получение шаблона по ID *В РАБОТЕ
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/template/:id [get]
func (*TTemplate) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Template.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template GetWithIndicators	godoc
// @Summary	Получение шаблона с вопросами по ID *В РАБОТЕ
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.TemplateWithIndicators
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/template/with-indicators/:id [get]
func (*TTemplate) GetWithIndicators(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Template.GetWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template GetAll	godoc
// @Summary	Получение всех шаблонов *В РАБОТЕ
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/template/all [get]
func (*TTemplate) GetAll(context *gin.Context) {
	data, errData := service.Template.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Template Create	godoc
// @Summary	Создание шаблона
// @Tags	template
// @Accept	json
// @Produce	json
// @Param request body models.TemplateCreateRequest true "Body для создания шаблона"
// @Success	200	{object}	models.TemplateWithCountIndicators
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/template/create [post]
func (*TTemplate) Create(context *gin.Context) {
	body := &models.TemplateCreateRequest{}
	utils.CultivateBody(context, body)

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
	isServiceError := utils.CultivateServiceError(context, errData)
	if isServiceError {
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
// @Summary	Скрытие шаблона по ID
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/template/hide/{id} [put]
func (*TTemplate) Hide(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template Delete	godoc
// @Summary	Удаление шаблона по ID
// @Tags	template
// @Accept	json
// @Produce	json
// @Param	id path int true "ID шаблона"
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/template/delete/{id} [delete]
func (*TTemplate) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
