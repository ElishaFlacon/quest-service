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
// @Success	200	{object}	models.QuestResponse
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/{id} [get]
func (*TQuest) Get(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.Get(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithIndicators	godoc
// @Summary	Получение опроса c его вопросами по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	models.QuestWithIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/with-indicators/{id} [get]
func (*TQuest) GetWithIndicators(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.GetWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsers	godoc
// @Summary	Получение опроса c его пользователями по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	models.QuestWithUsers
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/with-users/{id} [get]
func (*TQuest) GetWithUsers(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.GetWithUsers(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetWithUsersAndIndicators	godoc
// @Summary	Получение опроса c его вопросами и пользователями по ID
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	models.QuestWithUsersAndIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/with-users-and-indicators/{id} [get]
func (*TQuest) GetWithUsersAndIndicators(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.GetWithUsersAndIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetByUserId	godoc
// @Summary	Получение опросов по ID пользователя
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path string true "ID пользователя"
// @Success	200	{array}	models.QuestWithIndicators
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/by-user/with-indicators/{id} [get]
func (*TQuest) GetByUserIdWithIndicators(context *gin.Context) {
	id, errParam := utils.CultivateStringParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.GetByUserIdWithIndicators(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetByUserIdWithStatuses	godoc
// @Summary	Получение опросов по ID пользователя со статусами
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path string true "ID пользователя"
// @Success	200	{array}	models.QuestWithStatusesForUser
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/by-user/with-statuses/{id} [get]
func (*TQuest) GetByUserIdWithStatuses(context *gin.Context) {
	id, errParam := utils.CultivateStringParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.GetByUserIdWithStatuses(id)
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetAll	godoc
// @Summary	Получение всех опросов
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.QuestResponse
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/all [get]
func (*TQuest) GetAll(context *gin.Context) {
	data, errData := service.Quest.GetAll()
	utils.CultivateServiceData(context, data, errData)
}

// Quest GetAll	godoc
// @Summary	Получение всех опросов со статусами
// @Tags	quest
// @Accept	json
// @Produce	json
// @Success	200	{array}	models.QuestWithStatuses
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/all/with-statuses [get]
func (*TQuest) GetAllWithStatuses(context *gin.Context) {
	data, errData := service.Quest.GetAllWithStatuses()
	utils.CultivateServiceData(context, data, errData)
}

// Quest Create	godoc
// @Summary	Создание опроса
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param request body models.QuestCreate true "Body для создания опроса"
// @Param	Authorization header string true "Access token (с биркой)"
// @Success	200	{object}	models.Quest
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/create [post]
func (*TQuest) Create(context *gin.Context) {
	bearer := utils.GetBearer(context)

	body := &models.QuestCreate{}
	errBody := utils.CultivateBody(context, body)
	if errBody != nil {
		return
	}

	idTeams := utils.GetBodyIds(body.Teams)

	data, errData := service.Quest.Create(
		bearer,
		body.IdTemplate,
		idTeams,
		body.Name,
		body.Description,
		body.StartAt,
		body.EndAt,
	)
	utils.CultivateServiceData(context, data, errData)
}

// Quest Hide	godoc
// @Summary	Скрытие опроса по ID (используйте как удаление)
// @Tags	quest
// @Accept	json
// @Produce	json
// @Param	id path int true "ID опроса"
// @Success	200	{object}	models.Quest
// @Failure	400	{object} 	models.Error
// @Failure	500	{object} 	models.Error
// @Router	/api/v1/quest-service/quest/hide/{id} [put]
func (*TQuest) Hide(context *gin.Context) {
	id, errParam := utils.CultivateNumberParam(context, "id")
	if errParam != nil {
		return
	}

	data, errData := service.Quest.Hide(id)
	utils.CultivateServiceData(context, data, errData)
}
