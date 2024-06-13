package handle

import (
	"log"
	users "mymode/storage/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReqUser struct {
	Db *users.NewUserRepo
}

func (U *ReqUser) ConnectorDb(db *users.NewUserRepo) {
	U.Db = db
}

func (U *ReqUser) Create(c *gin.Context) {
	user := users.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	U.Db.Create(user)
	c.JSON(200, "Succes")
}

func (U *ReqUser) Read(c *gin.Context) {
	id := c.Param("id")
	user, err := U.Db.Read(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, user)
}

func (U *ReqUser) Update(c *gin.Context) {
	user := users.User{}

	id := c.Param("id")

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = U.Db.Update(user, id)
	if err != nil {
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, "Succes")
}

func (U *ReqUser) Delete(c *gin.Context) {
	id := c.Param("id")
	err := U.Db.Delete(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, "Succes")
}

func (U *ReqUser) GetAll(c *gin.Context) {
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
	users, err := U.Db.GetAll(lim, off)
	if err != nil{
		log.Fatal(err)
		c.JSON(404, err)
		return 
	}
	c.JSON(200, users)
}
