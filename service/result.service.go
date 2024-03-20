package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
)

type TResult struct{}

var Result *TResult

func (*TResult) Get() ([]models.Result, error) {
	data, err := database.Result.Get()
	return utils.NormalizeData(data, err)
}

func (*TResult) GetAll() ([]models.Result, error) {
	data, err := database.Result.GetAll()
	return utils.NormalizeData(data, err)
}

func (*TResult) Create(value string) error {
	err := database.Result.Create(value)
	return utils.NormalizeError(err)
}

func (*TResult) Update(value string) error {
	err := database.Result.Update(value)
	return utils.NormalizeError(err)
}

func (*TResult) Delete() error {
	err := database.Result.Delete()
	return utils.NormalizeError(err)
}

func (*TResult) DeleteAll() error {
	err := database.Result.DeleteAll()
	return utils.NormalizeError(err)
}
