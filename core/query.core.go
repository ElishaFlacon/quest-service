package core

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TQuery func(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
type TCopyFrom func(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)

// подходит для селектов
func QueryWithReturningData[T comparable](
	sqlString string,
	query TQuery,
	args ...any,
) ([]*T, error) {
	rows, err := query(context.Background(), sqlString, args...)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])

	if err != nil {
		return nil, err
	}

	return data, err
}

// подходит для инсертов, обновлений, удалений
func QueryWithReturningCount(
	tableName string,
	columnNames []string,
	rows [][]any,
	copyFrom TCopyFrom,
) (int64, error) {
	count, err := copyFrom(
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
