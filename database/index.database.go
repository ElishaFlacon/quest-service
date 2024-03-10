package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TInit struct {
	pool *pgxpool.Pool
}

func Init(database_url string) *TInit {
	pool, err := pgxpool.New(context.Background(), database_url)

	if err != nil {
		log.Fatal("Error create connection pool: ", database_url, err)
	}

	return &TInit{pool: pool}
}
