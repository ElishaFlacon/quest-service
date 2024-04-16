package controllers

// Деталка
// Создание

import (
	"github.com/gin-gonic/gin"
)

type TResult struct{}

var Result *TResult

// Result GetByUserId	godoc
// @Summary	Пример get result by user id
// @Tags	result
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Result
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/result/user/:id [get]
func (*TResult) GetByUserId(context *gin.Context) {
	// TODO
}

// Result GetByUsersId	godoc
// @Summary	Пример get result by users id
// @Tags	result
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Result
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/result/users [get]
func (*TResult) GetByUsersId(context *gin.Context) {
	// TODO
}

// Result Create	godoc
// @Summary	Пример result create
// @Tags	result
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Result
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/result/create [post]
func (*TResult) Create(context *gin.Context) {
	// TODO
}
