package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	port     = 5432
	dbname   = "newbase"
	password = "hamidjon4424"
)

func Connect() (*sql.DB, error) {
	connect := fmt.Sprintf("user = %s port = %d dbname = %s password = %s sslmode = disable", user, port, dbname, password)

	db, err := sql.Open("postgres", connect)

	return db, err
}
