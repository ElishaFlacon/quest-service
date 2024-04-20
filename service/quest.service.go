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

func (*TQuest) Get(id int) (*models.Quest, error) {
	// TODO для Тимура: возвращаем все поля для опроса
	return nil, nil
}

func (*TQuest) GetWithIndicators(id int) (*models.QuestWithIndicators, error) {
	// TODO для Тимура: возвращаем все поля для опроса + массив вопросов (use indicators.GetByQuestId)
	return nil, nil
}

func (*TQuest) GetAll() ([]*models.QuestResponse, error) {
	sqlString := `SELECT * FROM "launch_quests";`

	data, errData := database.BaseQuery[models.Quest](sqlString)
	if errData != nil {
		return nil, errData
	}

	quests := []*models.QuestResponse{}
	for index := range data {
		quest := data[index]

		// TODO доделать percent когда-нибудь
		percent := float32(0)
		status := utils.GetQuestTimeStatus(
			quest.StartAt,
			quest.EndAt,
		)

		newQuest := &models.QuestResponse{
			IdQuest: quest.IdQuest,
			Name:    quest.Name,
			StartAt: quest.StartAt,
			EndAt:   quest.EndAt,
			Percent: percent,
			Status:  status,
		}

		quests = append(quests, newQuest)
	}

	return quests, errData
}

func (*TQuest) Create(
	name string,
	description string,
	teams []string,
) (*models.Template, error) {
	// TODO продумать + доделать
	return nil, nil
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
