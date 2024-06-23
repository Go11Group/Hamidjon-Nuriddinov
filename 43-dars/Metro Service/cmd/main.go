package main

import (
	"mymode/api"
	"mymode/storage"
)

func main(){
	db, err := storage.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	panic(api.Router(db).Run(":8082"))
}