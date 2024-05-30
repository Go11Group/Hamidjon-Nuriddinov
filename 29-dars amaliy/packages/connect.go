package packages

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "mybase"
	password = "hamidjon4424"
)

func Connect()(*sql.DB, error){
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", connector)
	if err != nil{
		return nil, err
	}
	err = db.Ping()
	if err != nil{
		return nil, err
	}
	return db, nil
}

