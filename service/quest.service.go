package service

import (
	"errors"

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

func mapQuestTeamUserToString(
	questTeamUser *models.QuestTeamUser,
) string {
	return questTeamUser.IdUser
}

func getUsersData() ([]*models.UserWithStatusFull, error) {
	sqlString := `
		SELECT
			"quest".id_quest,
			"quest_team".id_team,
			"quest_team_user".id_user,
			"quest_team_user".name,
			"quest_team_user".email,
			(SELECT get_user_status(
				"quest".id_quest, 
				"quest_team_user".id_user
			)) status
		FROM "quest"
		INNER JOIN "quest_team" ON 
			"quest".id_quest = "quest_team".id_quest
		INNER JOIN "quest_team_user" ON 
			"quest_team".id_quest_team = "quest_team_user".id_quest_team
		WHERE "quest".available;
	`

	data, errData := database.BaseQuery[models.UserWithStatusFull](
		sqlString,
	)
	if errData != nil {
		return nil, errData
	}

	return data, nil
}

func getTeamsData() ([]*models.QuestTeamWithPercent, error) {
	sqlString := `
		SELECT *, (SELECT get_team_pass_percent(id_team, id_quest)) percent
		FROM "quest_team";
	`

	data, errData := database.BaseQuery[models.QuestTeamWithPercent](
		sqlString,
	)
	if errData != nil {
		return nil, errData
	}

	return data, nil
}

func (*TQuest) Get(id int) (*models.QuestResponse, error) {
	sqlString := `
		SELECT 
			id_quest,
			id_template,
			name,
			description,
			available,
			start_at,
			end_at, 
			(SELECT get_quest_pass_percent($1)) percent 
		FROM "quest" 
		WHERE id_quest = $1 AND available = true;
	`

	data, errData := database.BaseQuery[models.QuestWithPercent](
		sqlString,
		id,
	)
	if errData != nil {
		return nil, errData
	}

	quest := data[0]
	status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
	startAt := utils.GetStringDate(quest.StartAt)
	endAt := utils.GetStringDate(quest.EndAt)

	newQuest := &models.QuestResponse{
		IdQuest:    quest.IdQuest,
		IdTemplate: quest.IdTemplate,
		Name:       quest.Name,
		Percent:    quest.Percent,
		Available:  quest.Available,
		StartAt:    startAt,
		EndAt:      endAt,
		Status:     status,
	}

	return newQuest, nil
}

func (*TQuest) GetWithIndicators(id int) (*models.QuestWithIndicators, error) {
	quest, errQuest := Quest.Get(id)
	if errQuest != nil {
		return nil, errQuest
	}
	if quest == nil {
		return nil, errors.New("quest not found")
	}

	indicators, _ := Indicator.GetByQuestId(id)

	questWithIndicators := &models.QuestWithIndicators{
		IdQuest:    quest.IdQuest,
		IdTemplate: quest.IdTemplate,
		Name:       quest.Name,
		StartAt:    quest.StartAt,
		EndAt:      quest.EndAt,
		Status:     quest.Status,
		Percent:    quest.Percent,
		Available:  quest.Available,
		Indicators: indicators,
	}

	return questWithIndicators, nil
}

func (*TQuest) GetWithUsers(id int) (*models.QuestWithUsers, error) {
	quest, errQuest := Quest.Get(id)
	if errQuest != nil {
		return nil, errQuest
	}
	if quest == nil {
		return nil, errors.New("quest not found")
	}

	usersData, _ := QuestTeamUser.GetByQuestId(id)
	users := utils.MapToPrimitiveArray(
		usersData,
		mapQuestTeamUserToString,
	)

	questWithUsers := &models.QuestWithUsers{
		IdQuest:    quest.IdQuest,
		IdTemplate: quest.IdTemplate,
		Name:       quest.Name,
		StartAt:    quest.StartAt,
		EndAt:      quest.EndAt,
		Status:     quest.Status,
		Percent:    quest.Percent,
		Available:  quest.Available,
		Users:      users,
	}

	return questWithUsers, nil
}

func (*TQuest) GetWithUsersAndIndicators(
	id int,
) (*models.QuestWithUsersAndIndicators, error) {
	quest, errQuest := Quest.Get(id)
	if errQuest != nil {
		return nil, errQuest
	}
	if quest == nil {
		return nil, errors.New("quest not found")
	}

	indicators, _ := Indicator.GetByQuestId(id)
	usersData, _ := QuestTeamUser.GetByQuestId(id)
	users := utils.MapToPrimitiveArray(
		usersData,
		mapQuestTeamUserToString,
	)

	questWithUsersAndIndicators := &models.QuestWithUsersAndIndicators{
		IdQuest:    quest.IdQuest,
		IdTemplate: quest.IdTemplate,
		Name:       quest.Name,
		StartAt:    quest.StartAt,
		EndAt:      quest.EndAt,
		Status:     quest.Status,
		Percent:    quest.Percent,
		Available:  quest.Available,
		Users:      users,
		Indicators: indicators,
	}

	return questWithUsersAndIndicators, nil
}

func (*TQuest) GetByUserId(
	id string,
) ([]*models.QuestResponse, error) {
	sqlString := `
		SELECT
			"quest".id_quest,
			"quest".id_template,
			"quest".name,
			"quest".description,
			"quest".available,
			"quest".start_at,
			"quest".end_at,
			(SELECT get_quest_pass_percent("quest".id_quest)) percent
		FROM "quest" 
		INNER JOIN "quest_team" ON
			"quest".id_quest = "quest_team".id_quest
		INNER JOIN "quest_team_user" ON
			"quest_team".id_quest_team = "quest_team_user".id_quest_team
		WHERE "quest".available = true AND "quest_team_user".id_user = $1 AND "quest".available = true;
	`

	data, errData := database.BaseQuery[models.QuestWithPercent](
		sqlString,
		id,
	)
	if errData != nil {
		return nil, errData
	}

	result := []*models.QuestResponse{}

	for _, quest := range data {
		status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
		startAt := utils.GetStringDate(quest.StartAt)
		endAt := utils.GetStringDate(quest.EndAt)

		element := &models.QuestResponse{
			IdQuest:    quest.IdQuest,
			IdTemplate: quest.IdTemplate,
			Name:       quest.Name,
			Percent:    quest.Percent,
			Available:  quest.Available,
			StartAt:    startAt,
			EndAt:      endAt,
			Status:     status,
		}

		result = append(result, element)
	}

	return result, nil
}

func (*TQuest) GetByUserIdWithIndicators(
	id string,
) ([]*models.QuestWithIndicators, error) {
	data, errData := Quest.GetByUserId(id)
	if errData != nil {
		return nil, errData
	}

	result := []*models.QuestWithIndicators{}

	for _, quest := range data {
		// да, это супер не оптимизированно
		indicatorsData, _ := Indicator.GetByQuestId(quest.IdQuest)

		element := &models.QuestWithIndicators{
			IdQuest:    quest.IdQuest,
			IdTemplate: quest.IdTemplate,
			Name:       quest.Name,
			Percent:    quest.Percent,
			Available:  quest.Available,
			StartAt:    quest.StartAt,
			EndAt:      quest.EndAt,
			Status:     quest.Status,
			Indicators: indicatorsData,
		}

		result = append(result, element)
	}

	return result, nil
}

func (*TQuest) GetByUserIdWithStatuses(
	id string,
) ([]*models.QuestWithStatusesForUser, error) {
	usersData, errUsersData := getUsersData()
	if errUsersData != nil {
		return nil, errUsersData
	}

	teamsData, errTeamsData := getTeamsData()
	if errTeamsData != nil {
		return nil, errTeamsData
	}

	quests, errQuests := Quest.GetByUserId(id)
	if errQuests != nil {
		return nil, errQuests
	}

	// O(n^3) мы все виноваты в этом пиздеце
	questsWithUsers := []*models.QuestWithStatusesForUser{}
	for _, quest := range quests {
		isPass := false

		teams := []*models.TeamWithStatus{}
		for _, teamData := range teamsData {
			if quest.IdQuest != teamData.IdQuest {
				continue
			}

			users := []*models.UserWithStatus{}
			for _, userData := range usersData {
				isThisQuest := userData.IdQuest == quest.IdQuest
				isThisTeam := userData.IdTeam == teamData.IdTeam
				if !isThisQuest || !isThisTeam {
					continue
				}

				if *userData.IdUser == id {
					isPass = userData.Status
				}

				user := &models.UserWithStatus{
					IdUser: userData.IdUser,
					Name:   userData.Name,
					Email:  userData.Email,
					Status: userData.Status,
				}

				users = append(users, user)
			}

			team := &models.TeamWithStatus{
				IdTeam:  teamData.IdTeam,
				Name:    teamData.Name,
				Percent: teamData.Percent,
				Users:   users,
			}

			teams = append(teams, team)
		}

		element := &models.QuestWithStatusesForUser{
			IdQuest:    quest.IdQuest,
			IdTemplate: quest.IdTemplate,
			Name:       quest.Name,
			Percent:    quest.Percent,
			StartAt:    quest.StartAt,
			EndAt:      quest.EndAt,
			Status:     quest.Status,
			Available:  quest.Available,
			IsPass:     isPass,
			Teams:      teams,
		}

		questsWithUsers = append(questsWithUsers, element)
	}

	return questsWithUsers, nil
}

func (*TQuest) GetAll() ([]*models.QuestResponse, error) {
	sqlString := `
		SELECT 
			id_quest,
			id_template,
			name,
			description,
			available,
			start_at,
			end_at, 
			(SELECT get_quest_pass_percent(id_quest)) percent 
		FROM "quest" 
		WHERE available = true;
	`

	data, errData := database.BaseQuery[models.QuestWithPercent](sqlString)
	if errData != nil {
		return nil, errData
	}

	quests := []*models.QuestResponse{}
	for _, quest := range data {
		status := utils.GetQuestTimeStatus(quest.StartAt, quest.EndAt)
		startAt := utils.GetStringDate(quest.StartAt)
		endAt := utils.GetStringDate(quest.EndAt)

		element := &models.QuestResponse{
			IdQuest:    quest.IdQuest,
			IdTemplate: quest.IdTemplate,
			Name:       quest.Name,
			Percent:    quest.Percent,
			Available:  quest.Available,
			StartAt:    startAt,
			EndAt:      endAt,
			Status:     status,
		}

		quests = append(quests, element)
	}

	return quests, nil
}

func (*TQuest) GetAllWithStatuses() ([]*models.QuestWithStatuses, error) {
	usersData, errUsersData := getUsersData()
	if errUsersData != nil {
		return nil, errUsersData
	}

	teamsData, errTeamsData := getTeamsData()
	if errTeamsData != nil {
		return nil, errTeamsData
	}

	quests, errQuests := Quest.GetAll()
	if errQuests != nil {
		return nil, errQuests
	}

	// O(n^3) мы все виноваты в этом пиздеце
	questsWithUsers := []*models.QuestWithStatuses{}
	for _, quest := range quests {
		teams := []*models.TeamWithStatus{}
		for _, teamData := range teamsData {
			if quest.IdQuest != teamData.IdQuest {
				continue
			}

			users := []*models.UserWithStatus{}
			for _, userData := range usersData {
				isThisQuest := userData.IdQuest == quest.IdQuest
				isThisTeam := userData.IdTeam == teamData.IdTeam
				if !isThisQuest || !isThisTeam {
					continue
				}

				user := &models.UserWithStatus{
					IdUser: userData.IdUser,
					Name:   userData.Name,
					Email:  userData.Email,
					Status: userData.Status,
				}

				users = append(users, user)
			}

			team := &models.TeamWithStatus{
				IdTeam:  teamData.IdTeam,
				Name:    teamData.Name,
				Percent: teamData.Percent,
				Users:   users,
			}

			teams = append(teams, team)
		}

		element := &models.QuestWithStatuses{
			IdQuest:    quest.IdQuest,
			IdTemplate: quest.IdTemplate,
			Name:       quest.Name,
			Percent:    quest.Percent,
			StartAt:    quest.StartAt,
			EndAt:      quest.EndAt,
			Status:     quest.Status,
			Available:  quest.Available,
			Teams:      teams,
		}

		questsWithUsers = append(questsWithUsers, element)
	}

	return questsWithUsers, nil
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

	id := data[0].IdQuest
	link := utils.GetQuestLink(id)

	serviceTeams, errServiceTeams := Team.GetTeams(bearer, idTeams)
	if errServiceTeams != nil {
		Quest.Hide(id)
		return nil, errServiceTeams
	}

	questTeams, errQuestTeams := QuestTeam.CreateWithBatch(
		id,
		serviceTeams,
	)
	if errQuestTeams != nil {
		Quest.Hide(id)
		return nil, errQuestTeams
	}

	users := []*models.User{}
	teamUserArgs := []*models.QuestTeamUsers{}
	for _, questTeam := range questTeams {
		for _, serviceTeam := range serviceTeams {
			if questTeam.IdTeam != serviceTeam.IdTeam {
				continue
			}

			teamUserArg := &models.QuestTeamUsers{
				IdQuestTeam: questTeam.IdQuestTeam,
				Users:       serviceTeam.Users,
			}

			users = append(users, serviceTeam.Users...)
			teamUserArgs = append(teamUserArgs, teamUserArg)
		}
	}

	_, errTeamUser := QuestTeamUser.CreateWithBatch(
		teamUserArgs,
	)
	if errTeamUser != nil {
		Quest.Hide(id)
		return nil, errTeamUser
	}

	go Notification.SendNotification(users, link)

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

	return utils.CultivateFirstDataElement(data, err)
}
