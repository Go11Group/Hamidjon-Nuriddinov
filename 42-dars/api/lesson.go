package api

import (
	"log"
	"mymode/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReqLesson struct {
	Db Handle
}

func (L *ReqLesson) ConnectorDb(db Handle) {
	L.Db = db
}

func (L *ReqLesson) CreateLesson(c *gin.Context) {
	lesson := model.Lesson{}
	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	L.Db.l.CreateLesson(lesson)
	c.JSON(200, "Succes")
}

func (L *ReqLesson) ReadLesson(c *gin.Context) {
	id := c.Param("id")
	lesson, err := L.Db.l.ReadLesson(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, lesson)
}

func (L *ReqLesson) UpdateLesson(c *gin.Context) {
	lesson := model.Lesson{}

	id := c.Param("id")

	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = L.Db.l.UpdateLesson(lesson, id)
	if err != nil {
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, "Success")
}

func (L *ReqLesson) DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	err := L.Db.l.DeleteLesson(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, "Success")
}

func (L *ReqLesson) GetAllLessons(c *gin.Context) {
	params := make(map[string]interface{})
	filter := ""
	lessonFilter := model.LessonFilter{}
	var err error
	lessonFilter.CourseId = c.Query("course_id")
	lessonFilter.Title = c.Query("title")
	lessonFilter.Content = c.Query("content")
	limit := c.Query("limit")
	if limit != "" {
		lessonFilter.Limit, err = strconv.Atoi(limit)
		if err != nil {
			lessonFilter.Limit = 0
		}
	} else {
		lessonFilter.Limit = 0
	}
	offset := c.Query("offset")
	if offset != "" {
		lessonFilter.Offset, err = strconv.Atoi(offset)
		if err != nil {
			lessonFilter.Offset = 0
		}
	} else {
		lessonFilter.Offset = 0
	}

	if len(lessonFilter.CourseId) > 0 {
		params["course_id"] = lessonFilter.CourseId
		filter += " and course_id = :course_id"
	}

	if len(lessonFilter.Title) > 0 {
		params["title"] = lessonFilter.Title
		filter += " and title = :title"
	}

	if len(lessonFilter.Content) > 0 {
		params["content"] = lessonFilter.Content
		filter += " and content = :content"
	}

	if lessonFilter.Limit > 0 {
		params["limit"] = lessonFilter.Limit
		filter += " limit = :limit"
	}

	if lessonFilter.Offset > 0 {
		params["offset"] = lessonFilter.Offset
		filter += " offset = :offset"
	}

	query, arr := ReplaceQueryParamsLesson(filter, params)

	lessons, err := L.Db.l.GetAllLessons(query, arr)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, lessons)
}

func ReplaceQueryParamsLesson(query string, params map[string]interface{}) (string, []interface{}) {
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
