package service

import (
	"github.com/ElishaFlacon/quest-service/cruds"
	"github.com/ElishaFlacon/quest-service/models"
	"github.com/ElishaFlacon/quest-service/utils"
)

type TQuest struct {
	table string
}

var Quest = &TQuest{
	table: "quest",
}

func (*TQuest) GetAll() ([]models.QuestListResponse, error) {
	quests, err := cruds.LaunchQuest.GetAll()

	if err != nil {
		return nil, err
	}

	var questListResponse []models.QuestListResponse

	for index := range quests {
		quest := quests[index]

		questResponse := models.QuestListResponse{
			IdQuest: quest.IdQuest,
			Name:    quest.Name,
			StartAt: quest.StartAt,
			EndAt:   quest.EndAt,
			Percent: 12, // TODO ДОДЕЛАТЬ ПОДСЧЕТ ПРОЦЕНТА
			Status:  utils.GetTimeStatus(quest.StartAt, quest.EndAt),
		}

		questListResponse = append(questListResponse, questResponse)
	}

	return questListResponse, err

}
