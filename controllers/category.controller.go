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
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/category/{id} [get]
func (*TCategory) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
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
// @Failure	500	{string} 	string
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
// @Param request body models.CategoryCreateRequest true "Body для создания категории"
// @Success	200	{object}	models.Category
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/category/create [post]
func (*TCategory) Create(context *gin.Context) {
	body := models.CategoryCreateRequest{}
	utils.CultivateBody(context, body)
	data, errData := service.Category.Create(body.Name)
	utils.CultivateServiceData(context, data, errData)
}

// Category Delete	godoc
// @Summary	Удаление категории по ID
// @Tags	category
// @Accept	json
// @Produce	json
// @Param	id path int true "ID категории"
// @Success	200	{object}	models.Category
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/category/delete/{id} [delete]
func (*TCategory) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Category.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
