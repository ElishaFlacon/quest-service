package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuestIndicator struct{}

var QuestIndicator *TQuestIndicator

func (*TQuestIndicator) Get() ([]models.QuestIndicator, error) {
	data, err := database.QuestIndicator.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TQuestIndicator) GetAll() ([]models.QuestIndicator, error) {
	data, err := database.QuestIndicator.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TQuestIndicator) Create(value string) error {
	err := database.QuestIndicator.Create(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TQuestIndicator) Update(value string) error {
	err := database.QuestIndicator.Update(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TQuestIndicator) Delete() error {
	err := database.QuestIndicator.Delete()

	if err != nil {
		return err
	}

	return nil
}

func (*TQuestIndicator) DeleteAll() error {
	err := database.QuestIndicator.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
