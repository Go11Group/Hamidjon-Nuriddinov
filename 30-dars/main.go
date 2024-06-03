package main

import (
	"mymod/packages"
)

func main() {
	db, err := packages.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// u := packages.NewUserRepo(db)
	p := packages.NewProductRepo(db)

	// users := []packages.User{
	// 	{Username: "user1", Email: "user1@example.com", Password: "password1"},
	// 	{Username: "user2", Email: "user2@example.com", Password: "password2"},
	// 	{Username: "user3", Email: "user3@example.com", Password: "password3"},
	// 	{Username: "user4", Email: "user4@example.com", Password: "password4"},
	// 	{Username: "user5", Email: "user5@example.com", Password: "password5"},
	// 	{Username: "user6", Email: "user6@example.com", Password: "password6"},
	// 	{Username: "user7", Email: "user7@example.com", Password: "password7"},
	// 	{Username: "user8", Email: "user8@example.com", Password: "password8"},
	// 	{Username: "user9", Email: "user9@example.com", Password: "password9"},
	// 	{Username: "user10", Email: "user10@example.com", Password: "password10"},
	// }

	// for _, i := range users{
	// 	err := u.Create(i)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// }

	// users, err := u.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, i := range users {
	// 	fmt.Println(i)
	// }

	// user, err := u.ReadUser(15)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(user)

	// err = u.Update(packages.User{14, "hamidjon", "hamidjon@gmail.com", "1234"})
	// if err != nil{
	// 	panic(err)
	// }

	// err = u.Delete(12)
	// if err != nil{
	// 	panic(err)
	// }

	// products := []packages.Product{
	// 	{Id: 1, Name: "Product1", Description: "Description1", Price: 10.0, Stock_quantity: 100},
	// 	{Id: 2, Name: "Product2", Description: "Description2", Price: 20.0, Stock_quantity: 200},
	// 	{Id: 3, Name: "Product3", Description: "Description3", Price: 30.0, Stock_quantity: 300},
	// 	{Id: 4, Name: "Product4", Description: "Description4", Price: 40.0, Stock_quantity: 400},
	// 	{Id: 5, Name: "Product5", Description: "Description5", Price: 50.0, Stock_quantity: 500},
	// 	{Id: 6, Name: "Product6", Description: "Description6", Price: 60.0, Stock_quantity: 600},
	// 	{Id: 7, Name: "Product7", Description: "Description7", Price: 70.0, Stock_quantity: 700},
	// 	{Id: 8, Name: "Product8", Description: "Description8", Price: 80.0, Stock_quantity: 800},
	// 	{Id: 9, Name: "Product9", Description: "Description9", Price: 90.0, Stock_quantity: 900},
	// 	{Id: 10, Name: "Product10", Description: "Description10", Price: 100.0, Stock_quantity: 1000},
	// 	{Id: 11, Name: "Product11", Description: "Description11", Price: 110.0, Stock_quantity: 1100},
	// 	{Id: 12, Name: "Product12", Description: "Description12", Price: 120.0, Stock_quantity: 1200},
	// 	{Id: 13, Name: "Product13", Description: "Description13", Price: 130.0, Stock_quantity: 1300},
	// 	{Id: 14, Name: "Product14", Description: "Description14", Price: 140.0, Stock_quantity: 1400},
	// 	{Id: 15, Name: "Product15", Description: "Description15", Price: 150.0, Stock_quantity: 1500},
	// }

	// for _, i := range products{
	// 	err := p.Create(i)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// }

	// products, err := p.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, i := range products {
	// 	fmt.Println(i)
	// }

	// product, err := p.ReadUser(15)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(product)

	// err = p.Update(packages.Product{14, "Book", "very interisting", 50000, 3})
	// if err != nil{
	// 	panic(err)
	// }

	err = p.Delete(10)
	if err != nil{
		panic(err)
	}
}
