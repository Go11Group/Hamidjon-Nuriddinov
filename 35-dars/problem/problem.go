package problem

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type ProblemRepo struct {
	Db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{Db: db}
}

func (P *ProblemRepo) Create(w http.ResponseWriter, r *http.Request) {
	problem := Problem{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &problem)
	if err != nil {
		log.Fatal(err)
	}

	_, err = P.Db.Exec(`Insert Into Problem(status, name, difficulty) values ($1, $2, $3)`, problem.Status, problem.Name, problem.Difficulty)
	if err != nil {
		log.Fatal(err)
	}
}

func (P *ProblemRepo) Read(w http.ResponseWriter, r *http.Request){
	id := strings.Trim(r.URL.Path, "/problem/read/")
	problem := Problem{}

	err := P.Db.QueryRow(`Select status, name, difficulty from Problem where id = $1`, id).Scan(
		&problem.Status, &problem.Name, &problem.Difficulty)
	if err != nil{
		log.Fatal(err)
	}
	data, err := json.Marshal(problem)
	if err != nil{
		log.Fatal(err)
	}
	_, err = w.Write(data)
	if err != nil{
		log.Fatal(err)
	}
}

func (P *ProblemRepo) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.URL.Path, "/problem/update/")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	problem := Problem{}
	json.Unmarshal(body, &problem)

	_, err = P.Db.Exec(`Update Problem Set status = $1, name = $2, difficulty = $3 where id = $4`, 
problem.Status, problem.Name, problem.Difficulty, id)
	if err != nil{
		log.Fatal(err)
	}

}

func (P *ProblemRepo) Delete(w http.ResponseWriter, r *http.Request){
	id := strings.Trim(r.URL.Path, "/problem/update/")

	_, err := P.Db.Exec(`Delete from Problem where id = $1`, id)
	if err != nil{
		log.Fatal(err)
	}
}
