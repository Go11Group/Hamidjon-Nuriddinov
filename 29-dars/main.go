package main

import (
	"fmt"
	"mymod/packages"
)

func main() {
	db, err := packages.Connect()
	if err != nil {
		panic(err)
	}

	u := packages.NewUserRepo(db)

	// u.CreateTable()

	// u.Create(packages.User1{FirstName:"Hamidjon", LastName:"Nuriddinov", Email:"hamidjon@gmail.com", Password:"0409", Age:20, Field:"Gopher", Gender:"Male", IsEmployee:true})

	users := u.Read()
	for _, i := range users {
		fmt.Println(i)
	}

	//u.Update(packages.User1{FirstName:"Hamidjon", LastName:"Nuriddinov", Email:"hamidjon@gmail.com", Password:"0409", Age:20, Field:"Gopher", Gender:"Male", IsEmployee:true}, 2)
	// u.Delete(packages.User1{FirstName:"Hamidjon", LastName:"Nuriddinov", Email:"hamidjon@gmail.com", Password:"0409", Age:20, Field:"Gopher", Gender:"Male", IsEmployee:true})
}
