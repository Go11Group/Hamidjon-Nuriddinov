package product

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type ProductRepo struct {
	Db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{Db: db}
}

func (P *ProductRepo) Create(w http.ResponseWriter, r *http.Request) {
	product := Product{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Fatal(err)

	}
	_, err = P.Db.Exec("INSERT INTO Product(name, discription, price) values($1, $2, $3)", product.Name, product.Discription, product.Price)
	if err != nil {
		log.Fatal(err)

	}
}

func (P *ProductRepo) Read(w http.ResponseWriter, r *http.Request){
	product := Product{}
	id := strings.Trim(r.URL.Path, "/Product/read/")
	err := P.Db.QueryRow(`SELECT name, discription, price FROM Product WHERE id = $1`, 
	id).Scan(&product.Name, &product.Discription, &product.Price)
	if err != nil{
		log.Fatal(err)
	}
	data, err := json.Marshal(product)
	if err != nil{
		log.Fatal(err)
	}
	_, err = w.Write(data)
	if err != nil{
		log.Fatal(err)
	}
}


func (P *ProductRepo) Update(w http.ResponseWriter, r *http.Request){
	product := Product{}
	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &product)
	if err != nil{
		log.Fatal(err)
	}
	_, err = P.Db.Exec(`Update Product Set name = $1, discription = $2, price = $3 where id = $4`, 
		product.Name, product.Discription, product.Price, product.Id)
	if err != nil{
		log.Fatal(err)
	}
}

func (P *ProductRepo) Delete(w http.ResponseWriter, r *http.Request){
	id := strings.Trim(r.URL.Path, "/Product/delete/")
	_, err := P.Db.Exec(`Delete from Product where id = $1`, id)
	if err != nil{
		log.Fatal(err)
	}
}
