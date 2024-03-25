package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TQuest struct{}

var Quest *TQuest

func (*TQuest) Get(id int) (models.Quest, error) {
	data, err := core.CultivatingOneData(database.Quest.GetQuest(id))
	return *data, err
}

func (*TQuest) GetAll() ([]models.Quest, error) {
	data, err := database.Quest.GetAllQuests()
	return core.CultivatingData(data, err)
}

func (*TQuest) Create(value string) error {
	err := database.Quest.Create(value)
	return err
}

func (*TQuest) Update(value string) error {
	err := database.Quest.Update(value)
	return err
}

func (*TQuest) Delete() error {
	err := database.Quest.Delete()
	return err
}

func (*TQuest) DeleteAll() error {
	err := database.Quest.DeleteAll()
	return err
}
