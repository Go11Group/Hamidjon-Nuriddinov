package api

import (
	"log"
	"mymode/module"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqCourse struct {
	Db Handle
}

func (C *ReqCourse) ConnectorDb(db Handle) {
	C.Db = db
}

func (C *ReqCourse) CreateCourse(c *gin.Context) {
	course := module.Course{}
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
	course := module.Course{}

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
	courses, err := C.Db.c.GetAllCourses(lim, off)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, courses)
}
