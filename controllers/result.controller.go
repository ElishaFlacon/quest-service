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
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/result/by-user/{id} [get]
func (*TResult) GetByUserId(context *gin.Context) {
	id := utils.CultivateStringParam(context, "id")
	data, errData := service.Result.GetByUserId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Result GetByUsersId	godoc
// @Summary	Получение результатов по ID пользователей
// @Tags	result
// @Accept	json
// @Produce	json
// @Param request body models.GetByUsersIdRequest true "Body для получения результатов по ID пользователей"
// @Success	200	{array}	models.Result
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/result/by-users [get]
func (*TResult) GetByUsersId(context *gin.Context) {
	body := &models.GetByUsersIdRequest{}
	utils.CultivateBody(context, body)

	users := utils.GetBodyIds(body.Users)

	data, errData := service.Result.GetByUsersId(users)
	utils.CultivateServiceData(context, data, errData)
}

// Result Create	godoc
// @Summary	Создание результатов
// @Tags	result
// @Accept	json
// @Produce	json
// @Param request body models.ResultCreateRequest true "Body для создания результатов"
// @Success	200	{array}	models.Result
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/result/create [post]
func (*TResult) Create(context *gin.Context) {
	body := &models.ResultCreateRequest{}
	utils.CultivateBody(context, body)

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
	data := struct {
		CreatedResults int64 `json:"createdResults"`
	}{CreatedResults: count}
	utils.CultivateServiceData(context, data, errData)
}
