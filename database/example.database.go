package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IExample interface {
	Get() ([]*models.Example, error)
	Set(rows [][]any) (int64, error)
}

var Example IExample

func (repo *TDatabase) Get() ([]*models.Example, error) {
	sqlString := `SELECT * FROM example;`

	data, err := core.QueryWithReturningData[models.Example](sqlString, repo.db.Query)

	return data, err
}

func (repo *TDatabase) Set(rows [][]any) (int64, error) {
	tableName := "example"
	columnNames := []string{"value"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}
