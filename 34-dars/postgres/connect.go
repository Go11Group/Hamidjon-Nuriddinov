package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const(
	User = "postgres"
	Host = 5432
	Dbname = "newbase"
	Password = "hamidjon4424"
)

func Connect()(*sql.DB, error){
	connect := fmt.Sprintf("user = %s host = %d dbname = %s password = %s sslmode = disable", User, Host, Dbname, Password)

	db, err := sql.Open("postgres", connect)
	
	return db, err
}