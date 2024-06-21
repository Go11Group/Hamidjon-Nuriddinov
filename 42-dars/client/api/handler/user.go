package handler

import (
	"encoding/json"
	"fmt"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	resp, err := http.Get("http://localhost:8080/users/readUser/" + id)
	if err != nil {
		panic(err)
	}
	user := model.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	c.JSON(200, user)
}
