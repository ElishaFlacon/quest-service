package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TResult struct{}

var Result *TResult

func (*TResult) Get() ([]models.Result, error) {
	data, err := database.Result.Get()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TResult) GetAll() ([]models.Result, error) {
	data, err := database.Result.GetAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (*TResult) Create(value string) error {
	err := database.Result.Create(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TResult) Update(value string) error {
	err := database.Result.Update(value)

	if err != nil {
		return err
	}

	return nil
}

func (*TResult) Delete() error {
	err := database.Result.Delete()

	if err != nil {
		return err
	}

	return nil
}

func (*TResult) DeleteAll() error {
	err := database.Result.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
