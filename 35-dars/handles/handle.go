package handles

import (
	"database/sql"
	"mymode/problem"
	users "mymode/user"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(db *sql.DB) *http.Server {
	p := problem.NewProblemRepo(db)
	u := users.NewUserRepo(db)

	router := mux.NewRouter()

	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/create", u.Create).Methods("POST")
	user.HandleFunc("/read/", u.Read).Methods("GET")
	user.HandleFunc("/update/", u.Update).Methods("PUT")
	user.HandleFunc("/delete/", u.Delete).Methods("Delete")

	problem := router.PathPrefix("/problem/").Subrouter()
	problem.HandleFunc("/create", p.Create).Methods("POST")
	problem.HandleFunc("/read/", p.Read).Methods("GET")
	problem.HandleFunc("/update/", p.Update).Methods("PUT")
	problem.HandleFunc("/delete/", p.Delete).Methods("DELETE")


	return &http.Server{Addr: ":8080", Handler: router}
}
