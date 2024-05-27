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

func getQuestPercent(bearer string, id int) (int, error) {
	teams, errTeams := QuestTeam.GetByQuestId(id)
	if errTeams != nil {
		return 0, errTeams
	}

	idTeams := []string{}
	for _, team := range teams {
		idTeams = append(idTeams, team.IdTeam)
	}

	count, errCount := Team.GetUsersCount(bearer, idTeams)
	if errCount != nil {
		return 0, errTeams
	}

	users, errUsers := QuestTeamUser.GetByQuestId(id)
	if errUsers != nil {
		return 0, errUsers
	}

	percent := len(users) / count

	return percent, nil
}

func (*TQuest) Get(
	bearer string,
	id int,
) (*models.QuestResponse, error) {
	sqlString := `
		SELECT * FROM "quest" 
		WHERE id_quest = $1;
	`

	data, errData := database.BaseQuery[models.Quest](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	quest := data[0]
	percent, _ := getQuestPercent(bearer, id)
	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
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
	bearer string,
	id string,
) ([]*models.QuestWithIndicators, error) {
	sqlString := `
		SELECT
			"quest".id_quest,
			"quest".id_template,
			"quest".name,
			"quest".description,
			"quest".available,
			"quest".start_at,
			"quest".end_at
		FROM "quest" 
		INNER JOIN "quest_team" ON
			"quest".id_quest = "quest_team".id_quest
		INNER JOIN "quest_team_user" ON
			"quest_team_user".id_quest_team = "quest_team_user".id_quest_team
		WHERE "quest_team_user".id_user = $1;
	`

	data, errData := database.BaseQuery[models.Quest](sqlString, id)
	if errData != nil {
		return nil, errData
	}

	result := []*models.QuestWithIndicators{}

	for _, quest := range data {
		// да, это супер не оптимизированно
		indicatorsData, _ := Indicator.GetByQuestId(quest.IdQuest)
		percent, _ := getQuestPercent(bearer, quest.IdQuest)
		status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
		startAt := utils.GetStringDate(quest.StartAt)
		endAt := utils.GetStringDate(quest.EndAt)

		element := &models.QuestWithIndicators{
			IdQuest:    quest.IdQuest,
			Name:       quest.Name,
			StartAt:    startAt,
			EndAt:      endAt,
			Percent:    percent,
			Status:     status,
			Indicators: indicatorsData,
		}

		result = append(result, element)
	}

	return result, nil
}

func (*TQuest) GetWithIndicators(
	bearer string,
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

	indicatorsData, _ := Indicator.GetByQuestId(id)

	quest := data[0]
	percent, _ := getQuestPercent(bearer, id)
	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)

	newQuest := &models.QuestWithIndicators{
		IdQuest:    quest.IdQuest,
		Name:       quest.Name,
		StartAt:    startAt,
		EndAt:      endAt,
		Percent:    percent,
		Status:     status,
		Indicators: indicatorsData,
	}

	return newQuest, nil
}

func (*TQuest) GetWithUsers(
	bearer string,
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

	usersData, _ := QuestTeamUser.GetByQuestId(id)

	quest := data[0]
	percent, _ := getQuestPercent(bearer, id)
	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)
	users := utils.MapToPrimitiveArray(usersData, mapQuestTeamUserToString)

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
	bearer string,
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

	indicatorsData, _ := Indicator.GetByQuestId(id)
	usersData, _ := QuestTeamUser.GetByQuestId(id)

	quest := data[0]
	percent, _ := getQuestPercent(bearer, id)
	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)
	users := utils.MapToPrimitiveArray(usersData, mapQuestTeamUserToString)

	questWithUsersAndIndicators := &models.QuestWithUsersAndIndicators{
		IdQuest:    quest.IdQuest,
		Name:       quest.Name,
		StartAt:    startAt,
		EndAt:      endAt,
		Status:     status,
		Percent:    percent,
		Users:      users,
		Indicators: indicatorsData,
	}

	return questWithUsersAndIndicators, nil
}

func (*TQuest) GetAll(bearer string) ([]*models.QuestResponse, error) {
	sqlString := `SELECT * FROM "quest";`

	data, errData := database.BaseQuery[models.Quest](sqlString)
	if errData != nil {
		return nil, errData
	}

	quests := []*models.QuestResponse{}
	for index := range data {
		quest := data[index]

		// да, я знаю, это супер не оптимизировано
		percent, _ := getQuestPercent(bearer, quest.IdQuest)
		status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
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
	bearer string,
	idTemaplte int,
	idTeams []string,
	name string,
	description string,
	startAt int,
	endAt int,
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

	teams, errTeams := Team.GetTeams(bearer, idTeams)
	if errTeams != nil {
		return nil, errTeams
	}

	questTeamData, errQuestTeam := QuestTeam.CreateWithBatch(
		idQuest,
		idTeams,
	)
	if errQuestTeam != nil {
		return nil, errQuestTeam
	}

	teamUserArgs := []*models.QuestTeamUsers{}

	for index := range questTeamData {
		idQuestTeamUser := questTeamData[index].IdQuestTeam
		idTeamFromQuestTeam := questTeamData[index].IdTeam

		for indexTeam := range teams {
			idTeamFromMembers := teams[indexTeam].IdTeam
			users := teams[indexTeam].Users

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

	// в идеале тут должен был быть вызов notification service

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
