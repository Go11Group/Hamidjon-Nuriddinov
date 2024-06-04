package main

import (
	"database/sql"
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "newbase"
	password = "hamidjon4424"
)

func main() {
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", connector)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 1_000_000; i++{
		_, err = db.Exec("Insert Into Product values($1, $2, $3, $4)", uuid.New(), faker.Name(), faker.FirstName(), 10000)
		if err != nil{
			fmt.Println(err)
		}
		if i % 10000 == 0{
			fmt.Println(i)
		}
	}
}
