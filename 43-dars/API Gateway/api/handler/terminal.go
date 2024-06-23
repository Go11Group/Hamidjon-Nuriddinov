package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTerminal(c *gin.Context) {
	terminal := model.Terminal{}
	err := c.ShouldBindJSON(&terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := http.Post("http://localhost:8082/terminal/createTerminal", "JSON", bytes.NewBuffer(data))
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

func GetByIdTerminal(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8082/terminal/getByIdTerminal/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	terminal := model.Terminal{}
	err = json.NewDecoder(resp.Body).Decode(&terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, terminal)
}

func GetAllTerminals(c *gin.Context) {
	resp, err := http.Get("http://localhost:8082/terminal/getAllTerminals")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	terminals := []model.Terminal{}
	err = json.NewDecoder(resp.Body).Decode(&terminals)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, terminals)
}

func UpdateTerminal(c *gin.Context) {
	id := c.Param("id")
	terminal := model.Terminal{}
	err := c.ShouldBindJSON(&terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(terminal)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8082/terminal/updateTerminal/"+id, bytes.NewBuffer(data))
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

func DeleteTerminal(c *gin.Context) {
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8082/terminal/deleteTerminal/"+id, bytes.NewBuffer([]byte{}))
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
