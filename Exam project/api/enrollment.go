package api

import (
	"log"
	"mymode/module"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqEnrollment struct {
	Db Handle
}

func (E *ReqEnrollment) ConnectorDb(db Handle) {
	E.Db = db
}

func (E *ReqEnrollment) CreateEnrollment(c *gin.Context) {
	enrollment := module.Enrollment{}
	err := c.ShouldBindJSON(&enrollment)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = E.Db.e.CreateEnrollment(enrollment)
	if err != nil{
		log.Fatal(err)
		c.JSON(404, "Not found")
	}
	c.JSON(200, "Succes")
}

func (E *ReqEnrollment) ReadEnrollment(c *gin.Context) {
	id := c.Param("id")
	user, err := E.Db.e.ReadEnrollment(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, user)
}

func (E *ReqEnrollment) UpdateEnrollment(c *gin.Context) {
	enrollment := module.Enrollment{}

	id := c.Param("id")

	err := c.ShouldBindJSON(&enrollment)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = E.Db.e.UpdateEnrollment(enrollment, id)
	if err != nil {
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, "Succes")
}

func (E *ReqEnrollment) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")
	err := E.Db.e.DeleteEnrollment(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, "Success")
}

func (E *ReqEnrollment) GetAllEnrollments(c *gin.Context) {
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
	enrollments, err := E.Db.e.GetAllEnrollments(lim, off)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, enrollments)
}
