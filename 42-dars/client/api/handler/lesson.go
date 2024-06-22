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

func CreateLesson(c *gin.Context) {
	lesson := model.Lesson{}
	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	fmt.Println(lesson)
	data, err := json.Marshal(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8080/lessons/createLesson", "JSON", bytes.NewBuffer(data))
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

func ReadLesson(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8080/lessons/readLesson/" + id)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	lesson := model.Lesson{}
	err = json.NewDecoder(resp.Body).Decode(&lesson)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}

	c.JSON(200, lesson)
}

func GetAllLessons(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/lessons/getAllLessons")
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	lessons := []model.Lesson{}
	err = json.NewDecoder(resp.Body).Decode(&lessons)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	c.JSON(200, lessons)
}

func UpdateLesson(c *gin.Context) {
	lesson := model.Lesson{}
	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	data, err := json.Marshal(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	id := c.Param("id")

	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/lessons/updateLesson/"+id, bytes.NewBuffer(data))
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


func DeleteLesson(c *gin.Context){
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/lessons/deleteLesson/" + id, bytes.NewBuffer([]byte{}))
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
