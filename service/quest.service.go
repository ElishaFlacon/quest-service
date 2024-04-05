package service

import (
	"github.com/ElishaFlacon/questionnaire-service/cruds"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"time"
)

type TQuest struct {
	table string
}

//
//var Quest = &TQuest{
//	"quest",
//}

func getQuestStatusByStartAtAndEndAt(quest *models.Quest) string {
	currentTimeMinutes := time.Now().Minute()
	if quest.StartAt < currentTimeMinutes {
		return "Подготовлен"
	}
	if quest.StartAt >= currentTimeMinutes && quest.EndAt <= currentTimeMinutes {
		return "Запущен"
	}
	return "Завершен"

}

func (*TQuest) GetAll() ([]models.QuestListResponse, error) {
	quests, err := cruds.Quest.GetAll()
	if err == nil {
		var questListResponse []models.QuestListResponse

		for i := range quests {
			quest := quests[i]
			questResponse := models.QuestListResponse{
				IdQuest: quest.IdQuest,
				Name:    quest.Name,
				StartAt: quest.StartAt,
				EndAt:   quest.EndAt,
				Percent: 12, //пока не возможно посчитать
				Status:  getQuestStatusByStartAtAndEndAt(quest),
			}
			questListResponse = append(questListResponse, questResponse)
		}
		return questListResponse, err
	}
	return nil, err

}
