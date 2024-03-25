package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TExample struct{}

var Example *TExample

func (*TExample) GetAll() ([]*models.Example, error) {
	data, err := database.Example.GetAll()
	return core.CultivatingData(data, err)
}

func (*TExample) Create(rows [][]any) (int64, error) {
	count, err := database.Example.Create(rows)
	return count, err
}

func (*TExample) Update(id int, value string) ([]*models.Example, error) {
	data, err := database.Example.Update(id, value)
	return core.CultivatingData(data, err)
}
