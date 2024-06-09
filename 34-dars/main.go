package main

import (
	"mymode/handles"
	"mymode/postgres"
)

func main(){
	db, err := postgres.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()


	mux := handles.Handler(db)
	err = mux.ListenAndServe()
	if err != nil{
		panic(err)
	}
}