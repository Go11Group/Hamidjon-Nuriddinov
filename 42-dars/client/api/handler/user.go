package handler

import (
	"encoding/json"
	"log"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadUser(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8080/users/readUser/" + id)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	user := model.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}

	c.JSON(200, user)
}

func GetAll(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/users/getAllUsers")
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	users := []model.User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	c.JSON(200, users)
}
