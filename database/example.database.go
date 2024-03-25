package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IExample interface {
	Get(id int) ([]*models.Example, error)
	GetAll() ([]*models.Example, error)
	Create(rows [][]any) (int64, error)
	Update(id int, value string) ([]*models.Example, error)
	// Delete(id int) (int64, error)
	// DeleteAll() (int64, error)
}

var Example IExample

func (repo *TDatabase) Get(id int) ([]*models.Example, error) {
	sqlString := `SELECT * FROM "example" WHERE id=$1;`

	data, err := core.QueryWithReturningData[models.Example](sqlString, repo.db.Query, id)

	return data, err
}

func (repo *TDatabase) GetAll() ([]*models.Example, error) {
	sqlString := `SELECT * FROM "example";`

	data, err := core.QueryWithReturningData[models.Example](sqlString, repo.db.Query)

	return data, err
}

func (repo *TDatabase) Create(rows [][]any) (int64, error) {
	tableName := "example"
	columnNames := []string{"id", "value"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, repo.db.CopyFrom)

	return count, err
}

func (repo *TDatabase) Update(id int, value string) ([]*models.Example, error) {
	sqlString := `UPDATE "example" SET value=$2 WHERE id=$1 RETURNING *;`
	args := []any{id, value}

	data, err := core.QueryWithReturningData[models.Example](sqlString, repo.db.Query, args...)

	return data, err
}
