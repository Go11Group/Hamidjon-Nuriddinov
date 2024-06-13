package api

import (
	"log"
	"mymode/module"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqLesson struct {
	Db Handle
}

func (L *ReqLesson) ConnectorDb(db Handle) {
	L.Db = db
}

func (L *ReqLesson) CreateLesson(c *gin.Context) {
	lesson := module.Lesson{}
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
	lesson := module.Lesson{}

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
	var lim, off int
	limit := c.Param("limit")
	if limit == "" {
		lim = 0
	}
	lim, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, "Not found")
		return
	}

	offset := c.Param("offset")
	if offset == "" {
		off = 0
	}
	off, err = strconv.Atoi(offset)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, "Not found")
		return
	}
	lessons, err := L.Db.l.GetAllLessons(lim, off)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, lessons)
}
