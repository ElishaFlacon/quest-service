package utils

import (
	"fmt"
	"os"
)

func GetDatabaseString() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, name)
}
