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

func Init(database_url string) {
	db, err := pgxpool.New(context.Background(), database_url)

	if err != nil {
		log.Fatal("Error create connection pool: ", database_url, err)
	}

	Database = &TDatabase{db}
}

func BaseQuery[T comparable](sqlString string, args ...any) ([]*T, error) {
	rows, err := Database.Query(context.Background(), sqlString, args...)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])

	if err != nil {
		return nil, err
	}

	return data, err
}

// using for inserts
func CopyFromQuery(tableName string, columnNames []string, rows [][]any) (int64, error) {
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
