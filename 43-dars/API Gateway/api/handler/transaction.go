package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	transaction := model.Transaction{}
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := http.Post("http://localhost:8082/transaction/createTransaction", "JSON", bytes.NewBuffer(data))
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

func GetByIdTransaction(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8082/transaction/getByIdTransaction/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transaction := model.Transaction{}
	err = json.NewDecoder(resp.Body).Decode(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func GetAllTransactions(c *gin.Context) {
	resp, err := http.Get("http://localhost:8082/transaction/getAllTransactions")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transactions := []model.Transaction{}
	err = json.NewDecoder(resp.Body).Decode(&transactions)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction := model.Transaction{}
	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8082/transaction/updateTransaction/"+id, bytes.NewBuffer(data))
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

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8082/transaction/deleteTransaction/"+id, bytes.NewBuffer([]byte{}))
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
