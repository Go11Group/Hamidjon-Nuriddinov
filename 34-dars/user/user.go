package users

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (U *UserRepo) Create(w http.ResponseWriter, r *http.Request) {
	user := User{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)

	}
	_, err = U.Db.Exec("INSERT INTO Users(name, age) values($1, $2)", user.Name, user.Age)
	if err != nil {
		log.Fatal(err)

	}
}

func (U *UserRepo) Read(w http.ResponseWriter, r *http.Request) {
	user := User{}
	id := strings.Trim(r.URL.Path, "/User/read/")
	err := U.Db.QueryRow(`SELECT name, age FROM Users WHERE id = $1`,
		id).Scan(&user.Name, &user.Age)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func (U *UserRepo) Update(w http.ResponseWriter, r *http.Request) {
	user := User{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	_, err = U.Db.Exec(`Update Users Set name = $1, age = $2 where id = $4`, user.Name, user.Age, user.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func (U *UserRepo) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.URL.Path, "/User/delete/")
	_, err := U.Db.Exec(`Delete from Users where id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
}
