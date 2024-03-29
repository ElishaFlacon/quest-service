package database

import (
	"context"
	"log"

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
