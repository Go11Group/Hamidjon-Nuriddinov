package main

import "mymode/api"


func main(){
	r := api.Router()
	r.Run(":8081")
}