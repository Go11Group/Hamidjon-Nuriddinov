package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := http.Post("http://localhost:8081/user/createUser", "JSON", bytes.NewBuffer(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, string(body))
}

func GetByIdUser(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8081/user/getByIdUser/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user := model.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	resp, err := http.Get("http://localhost:8081/user/getAllUsers")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	users := []model.User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8081/user/updateUser/"+id, bytes.NewBuffer(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, string(body))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8081/user/deleteUser/"+id, bytes.NewBuffer([]byte{}))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, string(body))
}

func GetBalance(c *gin.Context){
	userId := c.Param("userId")
	cardId := c.Param("cardId")
	resp, err := http.Get("http://localhost:8081/user/getBalance/" + userId + "/" + cardId)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}
	balance := model.UserBalance{}
	err = json.NewDecoder(resp.Body).Decode(&balance)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, balance)
}
