package main

import (
	"mymode/api"
	"mymode/storage"
)

func main() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := api.Router(db)

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
