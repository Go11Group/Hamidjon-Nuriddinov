package main

import (
	"mymode/packages"
	"mymode/storage"

	"github.com/gin-gonic/gin"
)


func main() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	f := packages.NewRepo(db)

	r := gin.Default()

	r.GET("/getAll", f.GetAll)

	r.Run(":8080")

}
