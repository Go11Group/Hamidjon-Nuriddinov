package main

import (
	redisDB "homework/Redis"
	"homework/api"
	"log"
)

func main(){
	go redisDB.Server()
	
	router := api.Router()
	err := router.Run(":5555")
	if err != nil{
		log.Fatal(err)
	}
}