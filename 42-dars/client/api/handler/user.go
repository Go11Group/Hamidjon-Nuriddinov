package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	fmt.Println(user)
	data, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8080/users/createUser", "JSON", bytes.NewBuffer(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	fmt.Println(string(body))
	c.JSON(http.StatusCreated, string(body))
}

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

func GetAllUsers(c *gin.Context) {
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

func UpdateUser(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	data, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	id := c.Param("id")

	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/users/updateUser/"+id, bytes.NewBuffer(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, string(body))
}


func DeleteUser(c *gin.Context){
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/users/deleteUser/" + id, bytes.NewBuffer([]byte{}))
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, string(body))
}
