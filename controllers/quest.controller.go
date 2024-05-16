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
// @Summary	Получение опроса по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/{id} [get]
func (*TQuest) Get(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	bearer := utils.GetBearer(context)

	data, errData := service.Quest.Get(bearer, id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetByUserId	godoc
// @Summary	Получение опросов по ID пользователя
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID пользователя"
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{array}	models.QuestWithIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/by-user/{id} [get]
func (*TQuest) GetByUserId(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	bearer := utils.GetBearer(context)

	data, errData := service.Quest.GetByUserId(bearer, id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithIndicators	godoc
// @Summary	Получение опроса c его вопросами по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{object}	models.QuestWithIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/with-indicators/{id} [get]
func (*TQuest) GetWithIndicators(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	bearer := utils.GetBearer(context)

	data, errData := service.Quest.GetWithIndicators(bearer, id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsers	godoc
// @Summary	Получение опроса c его пользователями по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{object}	models.QuestWithUsers
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/with-users/{id} [get]
func (*TQuest) GetWithUsers(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	bearer := utils.GetBearer(context)

	data, errData := service.Quest.GetWithUsers(bearer, id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsersAndIndicators	godoc
// @Summary	Получение опроса c его вопросами и пользователями по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{object}	models.QuestWithUsersAndIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/with-users-and-indicators/{id} [get]
func (*TQuest) GetWithUsersAndIndicators(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	bearer := utils.GetBearer(context)

	data, errData := service.Quest.GetWithUsersAndIndicators(bearer, id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetAll	godoc
// @Summary	Получение всех опросов
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{array}	models.QuestResponse
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/all [get]
func (*TQuest) GetAll(context *gin.Context) {
	bearer := utils.GetBearer(context)
	data, errData := service.Quest.GetAll(bearer)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Create	godoc
// @Summary	Создание опроса
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param request body models.QuestCreateRequest true "Body для создания опроса"
// @Param	bearer-token header string true "Access token (с биркой)"
// @Success	200	{object}	models.Quest
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/create [post]
func (*TQuest) Create(context *gin.Context) {
	body := &models.QuestCreateRequest{}

	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	bearer := utils.GetBearer(context)
	IdTeams := utils.GetBodyIds(body.Teams)

	data, errData := service.Quest.Create(
		bearer,
		body.IdTemplate,
		IdTeams,
		body.Name,
		body.Description,
		body.StartAt,
		body.EndAt,
	)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Hide	godoc
// @Summary	Скрытие опроса по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	models.Quest
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/hide/{id} [put]
func (*TQuest) Hide(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Delete	godoc
// @Summary	Удаление опроса по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	models.Quest
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/quest-service/quest/delete/{id} [delete]
func (*TQuest) Delete(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.Delete(id)
	utils.CultivateServiceData(context, data, errData)
}
