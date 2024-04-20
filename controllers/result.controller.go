package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
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
	id := utils.CultivateStringParam(context, "id")
	data, errData := service.Result.GetByUserId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Result GetByUsersId	godoc
// @Summary	Пример get result by users ids
// @Tags	result
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.Result
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/result/users [get]
func (*TResult) GetByUsersId(context *gin.Context) {
	body := struct {
		Users []struct{ Id string } `json:"users"`
	}{}
	utils.CultivateBody(context, body)

	users := utils.GetBodyIds(body.Users)

	data, errData := service.Result.GetByUsersId(users)
	utils.CultivateServiceData(context, data, errData)
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
	body := struct {
		Results []struct {
			IdQuest     int    `json:"id_quest"`
			IdIndicator int    `json:"id_indicator"`
			IdFromUser  string `json:"id_from_user"`
			IdToUser    string `json:"id_to_user"`
			Value       string `json:"value"`
		} `json:"results"`
	}{}
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
