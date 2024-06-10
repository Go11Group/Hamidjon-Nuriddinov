package problem

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ProblemRepo struct{
	Db *sql.DB
}

func NewProblemRepo(db *sql.DB)*ProblemRepo{
	return &ProblemRepo{Db: db}
}

func (P *ProblemRepo) Create(w http.ResponseWriter, r *http.Request){
	problem := Problem{}

	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &problem)
	if err != nil{
		log.Fatal(err)
	}

	_, err = P.Db.Exec(`Insert Into Problem(status, name, difficulty) values ($1, $2, $3)`, problem.Status, problem.Name, problem.Difficulty)
	if err != nil{
		log.Fatal(err)
	}
}

