package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TIndicator struct{}

var Indicator *TIndicator

// Indicator Get	godoc
// @Summary	Пример get indicator by id
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/:id [get]
func (*TIndicator) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetAll	godoc
// @Summary	Пример get all indicators
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/:id [get]
func (*TIndicator) GetAll(context *gin.Context) {
	data, errData := service.Indicator.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetByTemplateId	godoc
// @Summary	Пример get indicator by template id
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/template/:id [get]
func (*TIndicator) GetByTemplateId(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.GetByTemplateId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator GetByQuestId	godoc
// @Summary	Пример get indicator by quest id
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/quest/:id [get]
func (*TIndicator) GetByQuestId(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.GetByQuestId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Create	godoc
// @Summary	Пример create indicator
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/create [post]
func (*TIndicator) Create(context *gin.Context) {
	body := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Role        string `json:"role"`
		Visible     bool   `json:"visible"`
		IdCategory  int    `json:"idCategory"`
	}{}
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
// @Summary	Пример hide indicator
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/hide/:id [put]
func (*TIndicator) Hide(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Indicator Delete	godoc
// @Summary	Пример delete indicator
// @Tags	indicator
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Indicator
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/indicator/delete/:id [delete]
func (*TIndicator) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Indicator.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
