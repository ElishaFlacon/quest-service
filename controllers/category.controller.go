package controllers

import (
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TCategory struct{}

var Category *TCategory

// Category Get	godoc
// @Summary	Получение категории по ID
// @Tags	category
// @Accept	json
// @Produce	json
// @Param	id path int true "ID категории"
// @Success	200	{object}	models.Category
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/category/{id} [get]
func (*TCategory) Get(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Category.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Category GetAll	godoc
// @Summary	Получение всех категорий
// @Tags	category
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Category
// @Failure	400	{string}	string
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/category/all [get]
func (*TCategory) GetAll(context *gin.Context) {
	data, errData := service.Category.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Category Create	godoc
// @Summary	Создание категории
// @Tags	category
// @Accept	json
// @Produce	json
// @Param request body models.CategoryCreate true "Body для создания категории"
// @Success	200	{object}	models.Category
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/category/create [post]
func (*TCategory) Create(context *gin.Context) {
	body := &models.CategoryCreate{}

	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	data, errData := service.Category.Create(body.Name)
	utils.CultivateServiceData(context, data, errData)
}

// Category Hide	godoc
// @Summary	Скрытие категории по ID (используйте как удаление)
// @Tags	category
// @Accept	json
// @Produce	json
// @Param	id path int true "ID категории"
// @Success	200	{object}	models.Category
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/category/hide/{id} [put]
func (*TCategory) Hide(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Category.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}
