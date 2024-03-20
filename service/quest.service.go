package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuest struct{}

var Quest *TQuest

func (*TQuest) Get() ([]models.Quest, error) {
	data, err := database.Quest.Get()
	return core.CultivatingData(data, err)
}

func (*TQuest) GetAll() ([]models.Quest, error) {
	data, err := database.Quest.GetAll()
	return core.CultivatingData(data, err)
}

func (*TQuest) Create(value string) error {
	err := database.Quest.Create(value)
	return core.CultivatingError(err)
}

func (*TQuest) Update(value string) error {
	err := database.Quest.Update(value)
	return core.CultivatingError(err)
}

func (*TQuest) Delete() error {
	err := database.Quest.Delete()
	return core.CultivatingError(err)
}

func (*TQuest) DeleteAll() error {
	err := database.Quest.DeleteAll()
	return core.CultivatingError(err)
}
