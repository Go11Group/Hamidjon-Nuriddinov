package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mymode/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStation(c *gin.Context) {
	station := model.Station{}
	err := c.ShouldBindJSON(&station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := http.Post("http://localhost:8082/station/createStation", "JSON", bytes.NewBuffer(data))
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

func GetByIdStation(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8082/station/getByIdStation/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	station := model.Station{}
	err = json.NewDecoder(resp.Body).Decode(&station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, station)
}

func GetAllStations(c *gin.Context) {
	resp, err := http.Get("http://localhost:8082/station/getAllStations")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	stations := []model.Station{}
	err = json.NewDecoder(resp.Body).Decode(&stations)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, stations)
}

func UpdateStation(c *gin.Context) {
	id := c.Param("id")
	station := model.Station{}
	err := c.ShouldBindJSON(&station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := json.Marshal(station)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8082/station/updateStation/"+id, bytes.NewBuffer(data))
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

func DeleteStation(c *gin.Context) {
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8082/station/deleteStation/"+id, bytes.NewBuffer([]byte{}))
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
