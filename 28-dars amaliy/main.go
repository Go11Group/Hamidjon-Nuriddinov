package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct{
	id int
	name string
	author string
	genre string
	price int 
	year int
}

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "mybase"
	password = "hamidjon4424"
)

func main(){
	db := Connect()
	defer db.Close()

	Delete(db)
}

func Connect()*sql.DB{
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", connector)
	if err != nil{
		panic(err)
	}

	err = db.Ping()
	if err != nil{
		panic(err)
	}
	return db
}


func Insert(db *sql.DB){
	var book Book
	fmt.Print("id = ")
	fmt.Scan(&book.id)
	fmt.Print("Name = ")
	fmt.Scan(&book.name)
	fmt.Print("Author = ")
	fmt.Scan(&book.author)
	fmt.Print("Genre = ")
	fmt.Scan(&book.genre)
	fmt.Print("Year = ")
	fmt.Scan(&book.year)
	fmt.Print("Price = ")
	fmt.Scan(&book.price)
	_, err := db.Exec("INSERT INTO kitoblar values ($1, $2, $3, $4, $5, $6)", book.id, book.name, book.author, book.genre, book.price, book.year)
	if err != nil{
		log.Fatal("Ma'lumot kiritishda xatolik!")
	}else{
		fmt.Println("Ma'lumot muvaffaqiyatli qo'shildi.")
	}
}

func Update(db *sql.DB){
	var book Book
	fmt.Println("Update qilmoqchi bo'lgan ma'lumotingizni idisini kiriting")
	fmt.Print("id = ")
	fmt.Scan(&book.id)
	fmt.Println("Yangi ma'lumotlarni kiriting")
	fmt.Print("Name = ")
	fmt.Scan(&book.name)
	fmt.Print("Author = ")
	fmt.Scan(&book.author)
	fmt.Print("Genre = ")
	fmt.Scan(&book.genre)
	fmt.Print("Year = ")
	fmt.Scan(&book.year)
	fmt.Print("Price = ")
	fmt.Scan(&book.price)

	_, err := db.Exec("Update kitoblar set name = $1, author = $2, genre = $3, year = $4, price = $5 where id = $6", book.name, book.author, book.genre, book.year, book.price, book.id)
	if err != nil{
		log.Fatal("Ma'lumotlar yangilanmadi!")
	}else{
		fmt.Println("Ma'lumotlar muvaffaqiyatli yangilandi.")
	}
}

func Delete(db *sql.DB){
	var id int
	fmt.Print("Id = ")
	fmt.Scan(&id)
	_, err := db.Exec("Delete from kitoblar where id = $1", id)
	if err != nil{
		log.Fatal("Ma'lumot o'chirilmadi")
	}else{
		fmt.Println("Ma'lumot muvaffaqiyatli o'chirildi.")
	}
}