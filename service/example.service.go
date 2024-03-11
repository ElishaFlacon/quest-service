package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TExample struct{}

var Example *TExample

func (*TExample) Get() ([]models.Example, error) {
	data, err := database.Example.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TExample) Set(value string) error {
	err := database.Example.Set(value)

	if err != nil {
		return err
	}

	return nil
}
