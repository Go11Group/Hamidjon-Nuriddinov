package api

import (
	"log"
	"mymode/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReqCourse struct {
	Db Handle
}

func (C *ReqCourse) ConnectorDb(db Handle) {
	C.Db = db
}

func (C *ReqCourse) CreateCourse(c *gin.Context) {
	course := model.Course{}
	err := c.ShouldBindJSON(&course)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	C.Db.c.CreateCourse(course)
	c.JSON(200, "Succes")
}

func (C *ReqCourse) ReadCourse(c *gin.Context) {
	id := c.Param("id")
	user, err := C.Db.c.ReadCourse(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, user)
}

func (C *ReqCourse) UpdateCourse(c *gin.Context) {
	course := model.Course{}

	id := c.Param("id")

	err := c.ShouldBindJSON(&course)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = C.Db.c.UpdateCourse(course, id)
	if err != nil {
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, "Success")
}

func (C *ReqCourse) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	err := C.Db.c.DeleteCourse(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, "Success")
}

func (C *ReqCourse) GetAllCourses(c *gin.Context) {
	params := make(map[string]interface{})
	filter := ""
	courseFilter := model.CourseFilter{}
	var err error
	courseFilter.Title = c.Query("title")
	courseFilter.Description = c.Query("description")
	limit := c.Query("limit")
	if limit != "" {
		courseFilter.Limit, err = strconv.Atoi(limit)
		if err != nil {
			courseFilter.Limit = 0
		}
	} else {
		courseFilter.Limit = 0
	}
	offset := c.Query("offset")
	if offset != "" {
		courseFilter.Offset, err = strconv.Atoi(offset)
		if err != nil {
			courseFilter.Offset = 0
		}
	} else {
		courseFilter.Offset = 0
	}

	if len(courseFilter.Title) > 0 {
		params["title"] = courseFilter.Title
		filter += " and title = :title"
	}

	if len(courseFilter.Description) > 0 {
		params["description"] = courseFilter.Description
		filter += " and description = :description"
	}

	if courseFilter.Limit > 0 {
		params["limit"] = courseFilter.Limit
		filter += " limit = :limit"
	}

	if courseFilter.Offset > 0 {
		params["offset"] = courseFilter.Offset
		filter += " offset = :offset"
	}

	query, arr := ReplaceQueryParamsCourse(filter, params)

	users, err := C.Db.c.GetAllCourses(query, arr)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, users)
}

func ReplaceQueryParamsCourse(query string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(query, ":"+k) {
			query = strings.ReplaceAll(query, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}
	return query, args
}

func (C *ReqCourse) GetLessonsByCourse(c *gin.Context) {
	id := c.Param("course_id")
	courseLessons, err := C.Db.c.GetLessonsByCourse(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, courseLessons)
}

func (C *ReqCourse) GetEnrolledUsersByCourse(c *gin.Context){
	course_id := c.Param("course_id")
	courseUsers, err := C.Db.c.GetEnrolledUsersByCourse(course_id)
	if err != nil{
		log.Fatal(err)
		c.JSON(404, "Not found")
	}
	c.JSON(200, courseUsers)
}
