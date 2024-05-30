package main

import (
	"fmt"
	"mymod/packages"
	users "mymod/packages/Users"
)

func main() {
	db, err := packages.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	u := users.NewUserRepo(db)

	// for _, i := range packages.ReadJson(){
	// 	err = u.Insert(i)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// }

	users, err := u.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, i := range users {
		fmt.Println(i)
	}

	// user, err := u.GetByID("062f4ebe-7844-47ec-a067-49f9f7601eae")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(user)
	
	// err = u.Update(users.User{"3f62e317-760b-4e18-81fd-e613baa6f6df", "Hamidjon", "Nuriddinov", "hamidjon@gmail.com", "Male", 20})
	// if err != nil{
	// 	panic(err)
	// }

	// err = u.Delete("d5af8087-4c70-480e-becf-aecab2e90eb0")
	// if err != nil{
	// 	panic(err)
	// }
}

