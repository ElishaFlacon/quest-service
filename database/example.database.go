package database

import (
	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/jackc/pgx/v5"
)

type IExample interface {
	Get() ([]models.Example, error)
	Set(value string) ([]models.Example, error)
}

var Example IExample

func (repo *TDatabase) Get() ([]models.Example, error) {
	sqlstring := `SELECT * FROM example;`

	cultivating := func(rows pgx.Rows) (models.Example, error) {
		var item models.Example

		err := rows.Scan(
			&item.Id,
			&item.Value,
		)

		return item, err
	}

	data, err := core.Query(
		sqlstring,
		cultivating,
		repo.db.Query,
	)

	return data, err
}

func (repo *TDatabase) Set(value string) ([]models.Example, error) {
	sqlstring := `INSERT INTO example (value) VALUES ($1);`

	cultivating := func(rows pgx.Rows) (models.Example, error) {
		var item models.Example

		err := rows.Scan(
			&item.Id,
			&item.Value,
		)

		return item, err
	}

	data, err := core.Query(
		sqlstring,
		cultivating,
		repo.db.Query,
		value,
	)

	return data, err

}
