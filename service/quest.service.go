package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuest struct{}

var Quest *TQuest

func (*TQuest) Get() ([]models.Quest, error) {
	data, err := database.Quest.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TQuest) GetAll() ([]models.Quest, error) {
	data, err := database.Quest.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TQuest) Create(value string) error {
	err := database.Quest.Create(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TQuest) Update(value string) error {
	err := database.Quest.Update(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TQuest) Delete() error {
	err := database.Quest.Delete()

	if err != nil {
		return err
	}

	return nil
}

func (*TQuest) DeleteAll() error {
	err := database.Quest.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
