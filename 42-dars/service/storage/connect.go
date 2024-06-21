package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	port     = 5432
	dbname   = "project"
	password = "hamidjon4424"
)

func Connect() (*sql.DB, error) {
	connector := fmt.Sprintf("user = %s port = %d dbname = %s password = %s sslmode = disable", user, port, dbname, password)

	db, err := sql.Open("postgres", connector)

	return db, err
}
