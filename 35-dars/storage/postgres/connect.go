package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	User     = "postgres"
	Port     = 5432
	Dbname   = "newbase"
	Password = "hamidjon4424"
)

func Connect()(*sql.DB, error){
	connector := fmt.Sprintf("user = %s port = %d dbname = %s password = %s sslmode = disable", User, Port, Dbname, Password)

	db, err := sql.Open("postgres", connector)

	return db, err
}
