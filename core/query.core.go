package core

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Query[T comparable](
	sql string,
	cultivating func(rows pgx.Rows) (T, error),
	query func(ctx context.Context, sql string, args ...any) (pgx.Rows, error),
	args ...any,
) ([]T, error) {
	rows, err := query(context.Background(), sql, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []T

	for rows.Next() {
		item, err := cultivating(rows)

		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}
