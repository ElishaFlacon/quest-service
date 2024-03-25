package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuestIndicator struct{}

var QuestIndicator *TQuestIndicator

func (*TQuestIndicator) Get(id int) (models.QuestIndicator, error) {
	data, err := core.CultivatingOneData(database.QuestIndicator.GetQuestIndicator(id))
	return *data, err
}

func (*TQuestIndicator) GetAll() ([]models.QuestIndicator, error) {
	data, err := database.QuestIndicator.GetAllQuestIndicators()
	return core.CultivatingData(data, err)
}

func (*TQuestIndicator) Create(value string) error {
	err := database.QuestIndicator.Create(value)
	return err
}

func (*TQuestIndicator) Update(value string) error {
	err := database.QuestIndicator.Update(value)
	return err
}

func (*TQuestIndicator) Delete() error {
	err := database.QuestIndicator.Delete()
	return err
}

func (*TQuestIndicator) DeleteAll() error {
	err := database.QuestIndicator.DeleteAll()
	return err
}
