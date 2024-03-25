package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IResult interface {
	GetResult(id int) ([]models.Result, error)
	GetAllResults() ([]models.Result, error)
	Create(value string) error
	Update(value string) error
	Delete() error
	DeleteAll() error
}

var Result IResult

func (repo *TDatabase) GetResult(id int) ([]*models.Result, error) {
	sqlString := `SELECT * FROM result WHERE idResult = $1;`

	data, err := core.QueryWithReturningData[models.Result](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAllResults() ([]*models.Result, error) {
	sqlString := `SELECT * FROM result;`

	data, err := core.QueryWithReturningData[models.Result](sqlString, repo.db.Query)

	return data, err
}
