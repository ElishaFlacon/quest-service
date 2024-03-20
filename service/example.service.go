package service

import (
	"github.com/ElishaFlacon/questionnaire-service/database"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/ElishaFlacon/questionnaire-service/utils"
)

type TExample struct{}

var Example *TExample

func (*TExample) Get() ([]models.Example, error) {
	data, err := database.Example.Get()
	return utils.NormalizeData(data, err)
}

func (*TExample) Set(value string) error {
	err := database.Example.Set(value)
	return utils.NormalizeError(err)
}
