package problem

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProblemRepo struct {
	Db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{Db: db}
}

func (P *ProblemRepo) Create(c *gin.Context) {
	problem := Problem{}

	body, err := io.ReadAll(c.Request.Body)
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

func (P *ProblemRepo) Read(c *gin.Context) {
	id := strings.Trim(c.Request.URL.Host, "/problem/read/")
	problem := Problem{}

	err := P.Db.QueryRow(`Select status, name, difficulty from Problem where id = $1`, id).Scan(
		&problem.Status, &problem.Name, &problem.Difficulty)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(problem)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func (P *ProblemRepo) Update(c *gin.Context) {
	id := strings.Trim(c.URL.Path, "/problem/update/")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	problem := Problem{}
	json.Unmarshal(body, &problem)

	_, err = P.Db.Exec(`Update Problem Set status = $1, name = $2, difficulty = $3 where id = $4`,
		problem.Status, problem.Name, problem.Difficulty, id)
	if err != nil {
		log.Fatal(err)
	}

}

func (P *ProblemRepo) Delete(c *gin.Context) {
	id := strings.Trim(r.URL.Path, "/problem/update/")

	_, err := P.Db.Exec(`Delete from Problem where id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
}
