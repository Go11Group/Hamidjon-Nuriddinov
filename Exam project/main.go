package main

import (
	"log"
	"mymode/api"
	"mymode/storage"
	users "mymode/storage/user"
)

func main() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := api.Router(users.NewUser(db))

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
