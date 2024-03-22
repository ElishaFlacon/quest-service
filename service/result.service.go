package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TResult struct{}

var Result *TResult

func (*TResult) Get() ([]models.Result, error) {
	data, err := database.Result.Get()
	return core.CultivatingData(data, err)
}

func (*TResult) GetAll() ([]models.Result, error) {
	data, err := database.Result.GetAll()
	return core.CultivatingData(data, err)
}

func (*TResult) Create(value string) error {
	err := database.Result.Create(value)
	return core.CultivatingError(err)
}

func (*TResult) Update(value string) error {
	err := database.Result.Update(value)
	return core.CultivatingError(err)
}

func (*TResult) Delete() error {
	err := database.Result.Delete()
	return core.CultivatingError(err)
}

func (*TResult) DeleteAll() error {
	err := database.Result.DeleteAll()
	return core.CultivatingError(err)
}
