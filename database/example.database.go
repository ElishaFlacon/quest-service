package database

import (
	"context"

	"github.com/ElishaFlacon/questionnaire-service/models"
)

type IExample interface {
	Get() ([]models.Example, error)
	Set(value string) error
}

var Example IExample

func (repo *TDatabase) Get() ([]models.Example, error) {
	rows, err := repo.db.Query(
		context.Background(),
		`SELECT * FROM example;`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Example

	for rows.Next() {
		var item models.Example

		err := rows.Scan(
			&item.Id,
			&item.Value,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}

func (repo *TDatabase) Set(value string) error {
	rows, err := repo.db.Query(
		context.Background(),
		`INSERT INTO example (value) VALUES ($1);`,
		value,
	)

	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
