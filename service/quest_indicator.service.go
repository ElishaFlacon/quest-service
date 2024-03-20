package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
)

type TQuestIndicator struct{}

var QuestIndicator *TQuestIndicator

func (*TQuestIndicator) Get() ([]models.QuestIndicator, error) {
	data, err := database.QuestIndicator.Get()
	return utils.NormalizeData(data, err)
}

func (*TQuestIndicator) GetAll() ([]models.QuestIndicator, error) {
	data, err := database.QuestIndicator.GetAll()
	return utils.NormalizeData(data, err)
}

func (*TQuestIndicator) Create(value string) error {
	err := database.QuestIndicator.Create(value)
	return utils.NormalizeError(err)
}

func (*TQuestIndicator) Update(value string) error {
	err := database.QuestIndicator.Update(value)
	return utils.NormalizeError(err)
}

func (*TQuestIndicator) Delete() error {
	err := database.QuestIndicator.Delete()
	return utils.NormalizeError(err)
}

func (*TQuestIndicator) DeleteAll() error {
	err := database.QuestIndicator.DeleteAll()
	return utils.NormalizeError(err)
}
