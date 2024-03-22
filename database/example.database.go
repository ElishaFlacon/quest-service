package database

import (
	"context"

	"github.com/ElishaFlacon/questionnaire-service/core"
	"github.com/ElishaFlacon/questionnaire-service/models"
	"github.com/jackc/pgx/v5"
)

type IExample interface {
	Get() ([]*models.Example, error)
	Set(rows [][]any) (int64, error)
}

var Example IExample

func (repo *TDatabase) Get() ([]*models.Example, error) {
	rows, err := repo.db.Query(context.Background(), "SELECT * FROM example;")

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Example])

	if err != nil {
		return nil, err
	}

	return data, err
}

func (repo *TDatabase) Set(rows [][]any) (int64, error) {
	count, err := repo.db.CopyFrom(
		context.Background(),
		pgx.Identifier{"example"},
		[]string{"value"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return 0, err
	}

	return count, err
}
