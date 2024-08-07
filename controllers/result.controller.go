package controllers

import (
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TResult struct{}

var Result *TResult

// Result GetByUserId	godoc
// @Summary	Получение результатов по ID пользователя
// @Tags	result
// @Accept	json
// @Produce	json
// @Param	id path string true "ID пользователя"
// @Success	200	{array}	models.Result
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/result/by-user/{id} [get]
func (*TResult) GetByUserId(context *gin.Context) {
	id, errParam := utils.CultivateStringParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Result.GetByUserId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Result GetByUsersId	godoc
// @Summary	Получение результатов по ID пользователей
// @Tags	result
// @Accept	json
// @Produce	json
// @Param request body models.UsersId true "Body для получения результатов по ID пользователей"
// @Success	200	{array}	models.Result
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/result/by-users [post]
func (*TResult) GetByUsersId(context *gin.Context) {
	body := &models.UsersId{}

	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	users := utils.GetBodyIds(body.Users)

	data, errData := service.Result.GetByUsersId(users)
	utils.CultivateServiceData(context, data, errData)
}

// Result Create	godoc
// @Summary	Создание результатов
// @Tags	result
// @Accept	json
// @Produce	json
// @Param request body models.ResultCreate true "Body для создания результатов"
// @Success	200	{object}	models.ResultResponse
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/result/create [post]
func (*TResult) Create(context *gin.Context) {
	body := &models.ResultCreate{}

	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	rows := [][]any{}
	for index := range body.Results {
		result := body.Results[index]

		element := []any{
			result.IdQuest,
			result.IdIndicator,
			result.IdFromUser,
			result.IdToUser,
			result.Value,
		}
		rows = append(rows, element)
	}

	count, errData := service.Result.CreateWithCopy(rows)
	data := models.ResultResponse{CreatedResults: count}
	utils.CultivateServiceData(context, data, errData)
}
