package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TIndicator struct{}

var Indicator *TIndicator

func (*TIndicator) Get() ([]models.Indicator, error) {
	data, err := database.Indicator.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TIndicator) GetAll() ([]models.Indicator, error) {
	data, err := database.Indicator.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TIndicator) Create(value string) error {
	err := database.Indicator.Create(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TIndicator) Update(value string) error {
	err := database.Indicator.Update(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TIndicator) Delete() error {
	err := database.Indicator.Delete()

	if err != nil {
		return err
	}

	return nil
}

func (*TIndicator) DeleteAll() error {
	err := database.Indicator.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
