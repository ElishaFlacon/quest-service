package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TLaunchQuest struct{}

var LaunchQuest *TLaunchQuest

func (*TLaunchQuest) Get() ([]models.LaunchQuest, error) {
	data, err := database.LaunchQuest.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TLaunchQuest) GetAll() ([]models.LaunchQuest, error) {
	data, err := database.LaunchQuest.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TLaunchQuest) Create(value string) error {
	err := database.LaunchQuest.Create(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TLaunchQuest) Update(value string) error {
	err := database.LaunchQuest.Update(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TLaunchQuest) Delete() error {
	err := database.LaunchQuest.Delete()

	if err != nil {
		return err
	}

	return nil
}

func (*TLaunchQuest) DeleteAll() error {
	err := database.LaunchQuest.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
