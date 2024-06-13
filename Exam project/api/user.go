package api

import (
	"log"
	"mymode/module"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqUser struct {
	Db Handle
}

func (U *ReqUser) ConnectorDb(db Handle) {
	U.Db = db
}

func (U *ReqUser) CreateUser(c *gin.Context) {
	user := module.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = U.Db.u.CreateUser(user)
	if err != nil{
		log.Fatal(err)
		c.JSON(404, "Not found")
		return 
	}
	c.JSON(200, "Succes")
}

func (U *ReqUser) ReadUser(c *gin.Context) {
	id := c.Param("id")
	user, err := U.Db.u.ReadUser(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, user)
}

func (U *ReqUser) UpdateUser(c *gin.Context) {
	user := module.User{}

	id := c.Param("id")

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = U.Db.u.UpdateUser(user, id)
	if err != nil {
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, "Succes")
}

func (U *ReqUser) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := U.Db.u.DeleteUser(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, "Success")
}

func (U *ReqUser) GetAllUsers(c *gin.Context) {
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
	users, err := U.Db.u.GetAllUsers(lim, off)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, users)
}
