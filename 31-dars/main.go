package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "newbase"
	password = "hamidjon4424"
)

func main(){
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", connector)
	if err != nil{
		panic(err)
	}
	defer db.Close()
}