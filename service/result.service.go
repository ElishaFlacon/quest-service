package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TResult struct{}

var Result *TResult

func (*TResult) Get(id int) (models.Result, error) {
	data, err := core.CultivatingOneData(database.Result.GetResult(id))
	return *data, err
}

func (*TResult) GetAll() ([]models.Result, error) {
	data, err := database.Result.GetAllResults()
	return core.CultivatingData(data, err)
}

func (*TResult) Create(value string) error {
	err := database.Result.Create(value)
	return err
}

func (*TResult) Update(value string) error {
	err := database.Result.Update(value)
	return err
}

func (*TResult) Delete() error {
	err := database.Result.Delete()
	return err
}

func (*TResult) DeleteAll() error {
	err := database.Result.DeleteAll()
	return err
}
