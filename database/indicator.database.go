package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IIndicator interface {
	GetIndicator(id int) ([]models.Indicator, error)
	GetAllIndicators() ([]models.Indicator, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var Indicator IIndicator

func (repo *TDatabase) GetIndicator(id int) ([]*models.Indicator, error) {
	sqlString := `SELECT * FROM indicator WHERE idIndicator = $1;`

	data, err := core.QueryWithReturningData[models.Indicator](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllIndicators() ([]*models.Indicator, error) {
	sqlString := `SELECT * FROM indicator;`

	data, err := core.QueryWithReturningData[models.Indicator](sqlString, repo.db.Query)

	return data, err
}
