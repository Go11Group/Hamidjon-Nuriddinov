package handles

import (
	"database/sql"
	"mymode/problem"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(db *sql.DB) *http.Server {
	p := problem.NewProblemRepo(db)

	router := mux.NewRouter()

	// user := router.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/create", ).Methods("POST")
	// user.HandleFunc("/update").Methods("POST")

	problem := router.PathPrefix("/problem/").Subrouter()
	problem.HandleFunc("/create", p.Create).Methods("POST")

	// solvedProblem := router.PathPrefix("/solvedProblem").Subrouter()

	return &http.Server{Addr: ":8080", Handler: router}
}
