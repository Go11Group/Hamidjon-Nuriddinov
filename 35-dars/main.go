package main

import (
	"mymode/handles"
	"mymode/storage/postgres"
)

func main(){
	db, err := postgres.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()


	server := handles.NewHandler(db)
	server.ListenAndServe()
}