package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCard(c *gin.Context) {
	card := model.Card{}
	err := c.ShouldBindJSON(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := http.Post("http://localhost:8082/card/createCard", "JSON", bytes.NewBuffer(data))
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

func GetByIdCard(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8082/card/getByIdCard/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	card := model.Card{}
	err = json.NewDecoder(resp.Body).Decode(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, card)
}

func GetAllCards(c *gin.Context) {
	resp, err := http.Get("http://localhost:8082/card/getAllCards")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	cards := []model.Card{}
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, cards)
}

func UpdateCard(c *gin.Context) {
	id := c.Param("id")
	card := model.Card{}
	err := c.ShouldBindJSON(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(card)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8082/card/updateCard/"+id, bytes.NewBuffer(data))
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

func DeleteCard(c *gin.Context) {
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8082/card/deleteCard/"+id, bytes.NewBuffer([]byte{}))
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
