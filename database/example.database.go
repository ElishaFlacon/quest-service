package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
)

type TExample struct{}

var Example *TExample

func (*TExample) Get(id int) ([]*models.Example, error) {
	sqlString := `SELECT * FROM "example" WHERE id=$1;`

	data, err := core.QueryWithReturningData[models.Example](sqlString, Database.Query, id)

	return data, err
}

func (*TExample) GetAll() ([]*models.Example, error) {
	sqlString := `SELECT * FROM "example";`

	data, err := core.QueryWithReturningData[models.Example](sqlString, Database.Query)

	return data, err
}

func (*TExample) Create(rows [][]any) (int64, error) {
	tableName := "example"
	columnNames := []string{"id", "value"}

	count, err := core.QueryWithReturningCount(tableName, columnNames, rows, Database.CopyFrom)

	return count, err
}

func (*TExample) Update(id int, value string) ([]*models.Example, error) {
	sqlString := `UPDATE "example" SET value=$2 WHERE id=$1 RETURNING *;`
	args := []any{id, value}

	data, err := core.QueryWithReturningData[models.Example](sqlString, Database.Query, args...)

	return data, err
}
