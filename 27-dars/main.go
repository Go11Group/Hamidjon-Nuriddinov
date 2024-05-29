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
	dbname = "mybase"
	password = "hamidjon4424"
)

type Worker struct{
	FISH string
 	JobID int
	Degree string
	Salary float64 
}

func main(){
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", connector)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	jobs := []string{"developer", "designer", "QA", "Doctor", "Teacher"}
	for _, i := range jobs{
		_, err = db.Exec("INSERT INTO jobs(name) values ($1)", i)
		if err != nil{
			panic(err)
		}
	}

	workers := []Worker{
        {FISH: "John Doe", JobID: 1, Degree: "Bachelor", Salary: 50000},
        {FISH: "Jane Smith", JobID: 2, Degree: "Master", Salary: 60000},
        {FISH: "Alice Johnson", JobID: 3, Degree: "PhD", Salary: 70000},
        {FISH: "Bob Brown", JobID: 4, Degree: "Associate", Salary: 40000},
        {FISH: "Charlie Davis", JobID: 5, Degree: "Bachelor", Salary: 55000},
        {FISH: "Dave Evans", JobID: 1, Degree: "Master", Salary: 65000},
        {FISH: "Eve Foster", JobID: 2, Degree: "PhD", Salary: 75000},
        {FISH: "Grace Hill", JobID: 3, Degree: "Bachelor", Salary: 52000},
        {FISH: "Hank Irving", JobID: 4, Degree: "Master", Salary: 62000},
        {FISH: "Ivy Jones", JobID: 5, Degree: "PhD", Salary: 72000},
    }

	for _, i := range workers{
		_, err = db.Exec("INSERT INTO workers(fish, job_id, degre, salary) values ($1, $2, $3, $4)", i.FISH, i.JobID, i.Degree, i.Salary)
		if err != nil{
			panic(err)
		}
	}

	var fish, degre, job string
	var salary float64  
	err = db.QueryRow("select W.FISH, J.name, W.degre, W.salary from Workers as W join Jobs as J on J.id = W.job_id;").Scan(&fish, &job, &degre, &salary)
	if err != nil{
		panic(err)
	}
	fmt.Println(fish, job, salary, degre)

	rows, err := db.Query(`
        SELECT w.id, w.FISH, w.job_id, w.degre, w.salary, j.name 
        FROM workers w 
        JOIN jobs j ON w.job_id = j.id
    `)
    if err != nil {
        panic(err)
    }
    defer rows.Close()


}