package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
)

type TIndicator struct{}

var Indicator *TIndicator

func (*TIndicator) Get() ([]models.Indicator, error) {
	data, err := database.Indicator.Get()
	return utils.NormalizeData(data, err)
}

func (*TIndicator) GetAll() ([]models.Indicator, error) {
	data, err := database.Indicator.GetAll()
	return utils.NormalizeData(data, err)
}

func (*TIndicator) Create(value string) error {
	err := database.Indicator.Create(value)
	return utils.NormalizeError(err)
}

func (*TIndicator) Update(value string) error {
	err := database.Indicator.Update(value)
	return utils.NormalizeError(err)
}

func (*TIndicator) Delete() error {
	err := database.Indicator.Delete()
	return utils.NormalizeError(err)
}

func (*TIndicator) DeleteAll() error {
	err := database.Indicator.DeleteAll()
	return utils.NormalizeError(err)
}
