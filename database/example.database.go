package database

import (
	"context"

	"github.com/ElishaFlacon/questionnaire-service/models"
)

func (init *TInit) GetExample() ([]models.Example, error) {
	rows, err := init.pool.Query(
		context.Background(),
		`SELECT * FROM Example;`,
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

func (init *TInit) SetExample(value string) error {
	rows, err := init.pool.Query(
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
