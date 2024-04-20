package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TCategory struct{}

var Category *TCategory

// Category Get	godoc
// @Summary	Пример get category by id
// @Tags	category
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Category
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/category/:id [get]
func (*TCategory) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Category.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Category Get All	godoc
// @Summary	Пример get all categorys
// @Tags	category
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Category
// @Failure	400	{string}	string
// @Failure	500	{string} 	string
// @Router	/category/all [get]
func (*TCategory) GetAll(context *gin.Context) {
	data, errData := service.Category.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Category Create	godoc
// @Summary	Пример create category
// @Tags	category
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Category
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/category/create [post]
func (*TCategory) Create(context *gin.Context) {
	body := struct {
		Name string `json:"name"`
	}{}
	utils.CultivateBody(context, body)
	data, errData := service.Category.Create(body.Name)
	utils.CultivateServiceData(context, data, errData)
}

// Category Delete	godoc
// @Summary	Пример delete category by id
// @Tags	category
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Category
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/category/delete/:id [delete]
func (*TCategory) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Category.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
