package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TLaunchQuest struct{}

var LaunchQuest *TLaunchQuest

func (*TLaunchQuest) Get(id int) (models.LaunchQuest, error) {
	data, err := core.CultivatingOneData(database.LaunchQuest.GetLaunchQuest(id))
	return *data, err
}

func (*TLaunchQuest) GetAll() ([]models.LaunchQuest, error) {
	data, err := database.LaunchQuest.GetAllLaunchQuests()
	return core.CultivatingData(data, err)
}

func (*TLaunchQuest) Create(value string) error {
	err := database.LaunchQuest.Create(value)
	return err
}

func (*TLaunchQuest) Update(value string) error {
	err := database.LaunchQuest.Update(value)
	return err
}

func (*TLaunchQuest) Delete() error {
	err := database.LaunchQuest.Delete()
	return err
}

func (*TLaunchQuest) DeleteAll() error {
	err := database.LaunchQuest.DeleteAll()
	return err
}
