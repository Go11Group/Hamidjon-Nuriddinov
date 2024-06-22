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

func CreateCourse(c *gin.Context) {
	course := model.Course{}
	err := c.ShouldBindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	fmt.Println(course)
	data, err := json.Marshal(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:8080/courses/createCourse", "JSON", bytes.NewBuffer(data))
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

func ReadCourse(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8080/courses/readCourse/" + id)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	course := model.Course{}
	err = json.NewDecoder(resp.Body).Decode(&course)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}

	c.JSON(200, course)
}

func GetAllCourses(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/courses/getAllCourses")
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	courses := []model.Course{}
	err = json.NewDecoder(resp.Body).Decode(&courses)
	if err != nil {
		c.JSON(404, "Not found")
		log.Fatal(err)
	}
	c.JSON(200, courses)
}

func UpdateCourse(c *gin.Context) {
	course := model.Course{}
	err := c.ShouldBindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	data, err := json.Marshal(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)
	}
	id := c.Param("id")

	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/courses/updateCourse/"+id, bytes.NewBuffer(data))
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


func DeleteCourse(c *gin.Context){
	id := c.Param("id")
	client := http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/courses/deleteCourse/" + id, bytes.NewBuffer([]byte{}))
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
