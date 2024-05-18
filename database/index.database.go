package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TDatabase struct {
	*pgxpool.Pool
}

var Database *TDatabase

func Init(databaseUrl string) {
	db, err := pgxpool.New(
		context.Background(),
		databaseUrl,
	)

	if err != nil {
		log.Fatal(
			"Error create connection pool: ",
			databaseUrl,
			err,
		)
	}

	Database = &TDatabase{db}
}

// Обычный запрос в базу
func BaseQuery[T any](
	sqlString string,
	args ...any,
) (
	[]*T,
	error,
) {
	rows, err := Database.Query(
		context.Background(),
		sqlString,
		args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(
		rows,
		pgx.RowToAddrOfStructByName[T],
	)
	if err != nil {
		return nil, err
	}

	return data, err
}

// Используйте для запросов в цикле
func SendBatch[T any](batch *pgx.Batch) ([]*T, error) {
	batchResult := Database.SendBatch(
		context.Background(),
		batch,
	)

	rows, err := batchResult.Query()
	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(
		rows,
		pgx.RowToAddrOfStructByName[T],
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Используйте для больших вставок данных (insert)
func CopyFromQuery(
	tableName string,
	columnNames []string,
	rows [][]any,
) (
	int64,
	error,
) {
	count, err := Database.CopyFrom(
		context.Background(),
		pgx.Identifier{tableName},
		columnNames,
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return 0, err
	}

	return count, err
}
