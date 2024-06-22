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

func CreateEnrollment(c *gin.Context) {
	enrollment := model.Enrollment{}
	err := c.ShouldBindJSON(&enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	fmt.Println(enrollment)
	data, err := json.Marshal(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8080/enrollments/createEnrollment", "JSON", bytes.NewBuffer(data))
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

func ReadEnrollment(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8080/enrollments/readEnrollment/" + id)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	enrollment := model.Enrollment{}
	err = json.NewDecoder(resp.Body).Decode(&enrollment)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}

	c.JSON(200, enrollment)
}

func GetAllEnrollments(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/enrollments/getAllEnrollments")
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	enrollments := []model.Enrollment{}
	err = json.NewDecoder(resp.Body).Decode(&enrollments)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	c.JSON(200, enrollments)
}

func UpdateEnrollment(c *gin.Context) {
	enrollment := model.Enrollment{}
	err := c.ShouldBindJSON(&enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	data, err := json.Marshal(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	id := c.Param("id")

	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/enrollments/updateEnrollment/"+id, bytes.NewBuffer(data))
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


func DeleteEnrollment(c *gin.Context){
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/enrollments/deleteEnrollment/" + id, bytes.NewBuffer([]byte{}))
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
