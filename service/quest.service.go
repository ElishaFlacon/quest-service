package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
)

type TQuest struct{}

var Quest *TQuest

func (*TQuest) Get() ([]models.Quest, error) {
	data, err := database.Quest.Get()
	return utils.NormalizeData(data, err)
}

func (*TQuest) GetAll() ([]models.Quest, error) {
	data, err := database.Quest.GetAll()
	return utils.NormalizeData(data, err)
}

func (*TQuest) Create(value string) error {
	err := database.Quest.Create(value)
	return utils.NormalizeError(err)
}

func (*TQuest) Update(value string) error {
	err := database.Quest.Update(value)
	return utils.NormalizeError(err)
}

func (*TQuest) Delete() error {
	err := database.Quest.Delete()
	return utils.NormalizeError(err)
}

func (*TQuest) DeleteAll() error {
	err := database.Quest.DeleteAll()
	return utils.NormalizeError(err)
}
