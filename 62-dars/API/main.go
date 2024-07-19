package main

import (
	"api/api"
	"log"
)


func main(){
	router := api.Router()
	err := router.Run(":5555")
	if err != nil{
		log.Fatal(err)
	}
}