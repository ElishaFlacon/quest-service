package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
)

type TLaunchQuest struct{}

var LaunchQuest *TLaunchQuest

func (*TLaunchQuest) Get() ([]models.LaunchQuest, error) {
	data, err := database.LaunchQuest.Get()
	return utils.NormalizeData(data, err)
}

func (*TLaunchQuest) GetAll() ([]models.LaunchQuest, error) {
	data, err := database.LaunchQuest.GetAll()
	return utils.NormalizeData(data, err)
}

func (*TLaunchQuest) Create(value string) error {
	err := database.LaunchQuest.Create(value)
	return utils.NormalizeError(err)
}

func (*TLaunchQuest) Update(value string) error {
	err := database.LaunchQuest.Update(value)
	return utils.NormalizeError(err)
}

func (*TLaunchQuest) Delete() error {
	err := database.LaunchQuest.Delete()
	return utils.NormalizeError(err)
}

func (*TLaunchQuest) DeleteAll() error {
	err := database.LaunchQuest.DeleteAll()
	return utils.NormalizeError(err)
}
