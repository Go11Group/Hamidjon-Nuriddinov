package main

import "mymode/api"

func main(){
	panic(api.Router().Run(":8080"))
}