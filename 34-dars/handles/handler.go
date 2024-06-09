package handles

import (
	"database/sql"
	"mymode/product"
	users "mymode/user"
	"net/http"
)

func Handler(db *sql.DB) *http.Server {
	p := product.NewProductRepo(db)
	u := users.NewUserRepo(db)
	mux := http.NewServeMux()
	mux.HandleFunc("POST/Product/create", p.Create)
	mux.HandleFunc("GET/Product/read/", p.Read)
	mux.HandleFunc("POST/Product/update", p.Update)
	mux.HandleFunc("GET/Product/delete/", p.Delete)

	mux.HandleFunc("POST/User/create", u.Create)
	mux.HandleFunc("GET/User/read/", u.Read)
	mux.HandleFunc("POST/User/update", u.Update)
	mux.HandleFunc("GET/User/delete/", u.Delete)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
