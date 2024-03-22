package service

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TExample struct{}

var Example *TExample

func (*TExample) Get() ([]*models.Example, error) {
	data, err := database.Example.Get()
	return core.CultivatingData(data, err)
}

func (*TExample) Set(rows [][]any) (int64, error) {
	count, err := database.Example.Set(rows)
	return count, err
}
