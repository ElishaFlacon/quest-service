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
// @Summary	Пример get template by id
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/template/:id [get]
func (*TTemplate) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Template.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template GetWithIndicators	godoc
// @Summary	Пример get template by id with indicators
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/template/:id [get]
func (*TTemplate) GetWithIndicators(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Template.GetWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template GetAll	godoc
// @Summary	Пример get all templates
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/template/:id [get]
func (*TTemplate) GetAll(context *gin.Context) {
	data, errData := service.Template.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Template Create	godoc
// @Summary	Пример create template
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.TemplateWithCountIndicators
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/template/create [post]
func (*TTemplate) Create(context *gin.Context) {
	body := struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Indicators  []struct{ Id int } `json:"indicators"`
	}{}
	utils.CultivateBody(context, body)

	indicators := utils.GetBodyIds(body.Indicators)
	utils.CultivateCondition(
		context,
		len(indicators) == 0,
		http.StatusBadRequest,
		"Indicators zero array",
	)

	data, count, errData := service.Template.Create(
		body.Name,
		body.Description,
		indicators,
	)
	utils.CultivateServiceError(context, errData)

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
// @Summary	Пример hide template by id
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/template/hide/:id [put]
func (*TTemplate) Hide(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Template Delete	godoc
// @Summary	Пример delete template by id
// @Tags	template
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Template
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/template/delete/:id [delete]
func (*TTemplate) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
