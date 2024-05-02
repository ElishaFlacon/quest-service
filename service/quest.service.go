package service

import (
	"github.com/ElishaFlacon/quest-service/database"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TQuest struct {
	table string
}

var Quest = &TQuest{
	table: "quest",
}

func mapQuestTeamUserToString(questTeamUser *models.QuestTeamUser) string {
	return questTeamUser.IdUser
}

func (*TQuest) Get(id int) (*models.QuestResponse, error) {
	sqlString := `
		SELECT * FROM "quest" 
		WHERE id_quest = $1;
	`

	data, errData := database.BaseQuery[models.Quest](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	quest := data[0]
	// TODO доделать percent когда будет готова quest_team_user.service GetByQuestId
	percent := float32(0)

	status := utils.GetQuestTimeStatus(
		quest.StartAt,
		quest.EndAt,
	)

	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)

	newQuest := &models.QuestResponse{
		IdQuest: quest.IdQuest,
		Name:    quest.Name,
		StartAt: startAt,
		EndAt:   endAt,
		Percent: percent,
		Status:  status,
	}

	return newQuest, nil
}

func (*TQuest) GetByUserId(
	id int,
) ([]*models.QuestWithIndicators, error) {
	// TODO будет похож на GetWithIndicators, только под конкретного пользователя
	return nil, nil
}

func (*TQuest) GetWithIndicators(
	id int,
) (*models.QuestWithIndicators, error) {
	sqlString := `
		SELECT * FROM "quest" 
		WHERE id_quest = $1;
	`

	data, errData := database.BaseQuery[models.Quest](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	indicators, errIndicators := Indicator.GetByQuestId(id)
	if errIndicators != nil {
		return nil, errIndicators
	}

	quest := data[0]
	// TODO доделать percent когда будет готова quest_team_user.service GetByQuestId
	percent := float32(0)

	status := utils.GetQuestTimeStatus(
		quest.StartAt,
		quest.EndAt,
	)

	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)

	newQuest := &models.QuestWithIndicators{
		IdQuest:    quest.IdQuest,
		Name:       quest.Name,
		StartAt:    startAt,
		EndAt:      endAt,
		Percent:    percent,
		Status:     status,
		Indicators: indicators,
	}

	return newQuest, nil
}

func (*TQuest) GetWithUsers(
	id int,
) (*models.QuestWithUsers, error) {
	sqlString := `
		SELECT * FROM "quest" 
		WHERE id_quest = $1;
	`

	data, errData := database.BaseQuery[models.Quest](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	questTeamUsers, errQuestTeamUsers := QuestTeamUser.GetByQuestId(id)
	if errQuestTeamUsers != nil {
		return nil, errQuestTeamUsers
	}

	quest := data[0]

	percent := float32(0)

	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)

	users := utils.MapToPrimitiveArray(questTeamUsers, mapQuestTeamUserToString)

	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)

	questWithUsers := &models.QuestWithUsers{
		IdQuest: quest.IdQuest,
		Name:    quest.Name,
		StartAt: startAt,
		EndAt:   endAt,
		Status:  status,
		Percent: percent,
		Users:   users,
	}

	return questWithUsers, nil
}

func (*TQuest) GetWithUsersAndIndicators(
	id int,
) (*models.QuestWithUsersAndIndicators, error) {
	sqlString := `
		SELECT * FROM "quest" 
		WHERE id_quest = $1;
	`

	data, errData := database.BaseQuery[models.Quest](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	questTeamUsers, errQuestTeamUsers := QuestTeamUser.GetByQuestId(id)
	if errQuestTeamUsers != nil {
		return nil, errQuestTeamUsers
	}

	indicators, errIndicators := Indicator.GetByQuestId(id)
	if errIndicators != nil {
		return nil, errIndicators
	}

	quest := data[0]

	percent := float32(0)

	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)

	users := utils.MapToPrimitiveArray(questTeamUsers, mapQuestTeamUserToString)

	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)

	questWithUsersAndIndicators := &models.QuestWithUsersAndIndicators{
		IdQuest:    quest.IdQuest,
		Name:       quest.Name,
		StartAt:    startAt,
		EndAt:      endAt,
		Status:     status,
		Percent:    percent,
		Users:      users,
		Indicators: indicators,
	}

	return questWithUsersAndIndicators, nil
}

func (*TQuest) GetAll() ([]*models.QuestResponse, error) {
	sqlString := `SELECT * FROM "quest";`

	data, errData := database.BaseQuery[models.Quest](sqlString)
	if errData != nil {
		return nil, errData
	}

	quests := []*models.QuestResponse{}
	for index := range data {
		quest := data[index]

		// TODO доделать percent когда будет готова quest_team_user.service GetByQuestId
		percent := float32(0)
		status := utils.GetQuestTimeStatus(
			quest.StartAt,
			quest.EndAt,
		)

		startAt := utils.GetStringDate(quest.StartAt)
		endAt := utils.GetStringDate(quest.EndAt)

		newQuest := &models.QuestResponse{
			IdQuest: quest.IdQuest,
			Name:    quest.Name,
			StartAt: startAt,
			EndAt:   endAt,
			Percent: percent,
			Status:  status,
		}

		quests = append(quests, newQuest)
	}

	return quests, errData
}

func (*TQuest) Create(
	idTemaplte int,
	name string,
	description string,
	startAt int,
	endAt int,
	teams []string,
) (*models.Quest, error) {
	sqlString := `
		INSERT INTO "quest" 
		(id_template, name, description, start_at, end_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *;
	`
	args := []any{idTemaplte, name, description, startAt, endAt}

	data, errData := database.BaseQuery[models.Quest](
		sqlString,
		args...,
	)
	if errData != nil {
		return nil, errData
	}

	idQuest := data[0].IdQuest

	members, errMembers := Team.GetAllMembersOfTeams(teams)
	if errMembers != nil {
		return nil, errMembers
	}

	// TODO позже удалить
	// members := []models.TeamMembers{{IdTeam: "300", IdUsers: []string{"12", "13"}}}

	questTeamData, errQuestTeam := QuestTeam.CreateWithBatch(
		idQuest,
		teams,
	)
	if errQuestTeam != nil {
		return nil, errQuestTeam
	}

	teamUserArgs := []*models.QuestTeamUsers{}

	for index := range questTeamData {
		idQuestTeamUser := questTeamData[index].IdQuestTeam
		idTeamFromQuestTeam := questTeamData[index].IdTeam

		for indexTeam := range members {
			idTeamFromMembers := members[indexTeam].IdTeam
			users := members[indexTeam].IdUsers

			if idTeamFromMembers != idTeamFromQuestTeam {
				continue
			}

			element := &models.QuestTeamUsers{
				IdQuestTeam: idQuestTeamUser,
				Users:       users,
			}

			teamUserArgs = append(teamUserArgs, element)
		}
	}

	_, errTeamUser := QuestTeamUser.CreateWithBatch(
		teamUserArgs,
	)
	if errTeamUser != nil {
		return nil, errTeamUser
	}

	// TODO вызов notification service добавить когда-нибудь

	return data[0], nil
}

func (*TQuest) Hide(id int) (*models.Quest, error) {
	sqlString := `
		UPDATE "quest" 
		SET available = false
		WHERE id_quest = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Quest](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}

func (*TQuest) Delete(id int) (*models.Quest, error) {
	sqlString := `
		DELETE FROM "quest" 
		WHERE id_quest = $1 
		RETURNING *;
	`

	data, err := database.BaseQuery[models.Quest](sqlString, id)

	return utils.CultivateFirstDataElemet(data, err)
}
