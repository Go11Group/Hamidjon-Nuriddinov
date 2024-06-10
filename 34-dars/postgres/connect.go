package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const(
	User = "postgres"
	port = 5432
	Dbname = "newbase"
	Password = "hamidjon4424"
)

func Connect()(*sql.DB, error){
	connect := fmt.Sprintf("user = %s port = %d dbname = %s password = %s sslmode = disable", User, port, Dbname, Password)

	db, err := sql.Open("postgres", connect)
	
	return db, err
}