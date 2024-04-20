package controllers

import (
	"github.com/ElishaFlacon/quest-service/service"
	"github.com/ElishaFlacon/quest-service/utils"
	"github.com/gin-gonic/gin"
)

type TQuest struct{}

var Quest *TQuest

// Quest Get	godoc
// @Summary	Пример get quest by id
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/:id [get]
func (*TQuest) Get(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithIndicators	godoc
// @Summary	Пример get quest by id with indicators
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/with-indicators/:id [get]
func (*TQuest) GetWithIndicators(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsers	godoc
// @Summary	Пример get quest by id with users
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/with-users/:id [get]
func (*TQuest) GetWithUsers(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetWithUsers(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsersAndIndicators	godoc
// @Summary	Пример get quest by id with users and indicators
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/with-users-and-indicators/:id [get]
func (*TQuest) GetWithUsersAndIndicators(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.GetWithUsersAndIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetAll	godoc
// @Summary	Пример get all quest
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.QuestResponse
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/all [get]
func (*TQuest) GetAll(context *gin.Context) {
	data, errData := service.Quest.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Quest Create	godoc
// @Summary	Пример create quest
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/create [post]
func (*TQuest) Create(context *gin.Context) {
	body := struct {
		Name        string                `json:"name"`
		Description string                `json:"description"`
		Teams       []struct{ Id string } `json:"teams"`
	}{}
	utils.CultivateBody(context, body)

	teams := utils.CultivateId(body.Teams)
	data, errData := service.Quest.Create(
		body.Name,
		body.Description,
		teams,
	)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Hide	godoc
// @Summary	Пример hide quest by id
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/hide/:id [put]
func (*TQuest) Hide(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Delete	godoc
// @Summary	Пример delete quest by id
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{object}	models.Quest
// @Failure	400	{string} 	string
// @Failure	500	{string} 	string
// @Router	/quest/delete/:id [delete]
func (*TQuest) Delete(context *gin.Context) {
	id := utils.CultivateNumberParam(context, "id")
	data, errData := service.Quest.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
