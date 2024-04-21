package controllers

import (
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TQuest struct{}

var Quest *TQuest

// Quest Get	godoc
// @Summary	Получение опроса по ID *В РАБОТЕ
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/{id} [get]
func (*TQuest) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetByUserId	godoc
// @Summary	Получение опросов по ID пользователя *В РАБОТЕ
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.QuestWithIndicators
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/by-user/{id} [get]
func (*TQuest) GetByUserId(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetByUserId(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithIndicators	godoc
// @Summary	Получение опроса c его вопросами по ID *В РАБОТЕ
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestWithIndicators
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/with-indicators/{id} [get]
func (*TQuest) GetWithIndicators(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsers	godoc
// @Summary	Получение опроса c его пользователями по ID *В РАБОТЕ
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestWithUsers
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/with-users/{id} [get]
func (*TQuest) GetWithUsers(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetWithUsers(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsersAndIndicators	godoc
// @Summary	Получение опроса c его вопросами и пользователями по ID *В РАБОТЕ
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestWithUsersAndIndicators
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/with-users-and-indicators/{id} [get]
func (*TQuest) GetWithUsersAndIndicators(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetWithUsersAndIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetAll	godoc
// @Summary	Получение всех опросов
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/all [get]
func (*TQuest) GetAll(context *gin.Context) {
	data, errData := service.Quest.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Quest Create	godoc
// @Summary	Создание опроса
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param request body models.QuestCreateRequest true "Body для создания опроса"
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/create [post]
func (*TQuest) Create(context *gin.Context) {
	body := models.QuestCreateRequest{}
	utils.CultivateBody(context, body)

	teams := utils.GetBodyIds(body.Teams)

	data, errData := service.Quest.Create(
		body.Name,
		body.Description,
		teams,
	)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Hide	godoc
// @Summary	Скрытие опроса по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/hide/{id} [put]
func (*TQuest) Hide(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Delete	godoc
// @Summary	Удаление опроса по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest-service/quest/delete/{id} [delete]
func (*TQuest) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
