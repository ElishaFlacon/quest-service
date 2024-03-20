package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TIndicator struct{}

var Indicator *TIndicator

func (*TIndicator) Get() ([]models.Indicator, error) {
	data, err := database.Indicator.Get()
	return core.CultivatingData(data, err)
}

func (*TIndicator) GetAll() ([]models.Indicator, error) {
	data, err := database.Indicator.GetAll()
	return core.CultivatingData(data, err)
}

func (*TIndicator) Create(value string) error {
	err := database.Indicator.Create(value)
	return core.CultivatingError(err)
}

func (*TIndicator) Update(value string) error {
	err := database.Indicator.Update(value)
	return core.CultivatingError(err)
}

func (*TIndicator) Delete() error {
	err := database.Indicator.Delete()
	return core.CultivatingError(err)
}

func (*TIndicator) DeleteAll() error {
	err := database.Indicator.DeleteAll()
	return core.CultivatingError(err)
}
