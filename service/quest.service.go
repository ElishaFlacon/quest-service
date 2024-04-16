package service

// TODO add all methods

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

func (*TQuest) GetAll() ([]*models.QuestListResponse, error) {
	data, errData := cruds.Quest.GetAll()

	if errData != nil {
		return nil, errData
	}

	quests := []*models.QuestListResponse{}

	for index := range data {
		quest := data[index]

		// TODO ДОДЕЛАТЬ ПОДСЧЕТ ПРОЦЕНТА
		percent := float32(0)
		status := utils.GetQuestTimeStatus(
			quest.StartAt,
			quest.EndAt,
		)

		newQuest := &models.QuestListResponse{
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
