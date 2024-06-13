package api

import (
	"mymode/api/handle"
	users "mymode/storage/user"

	"github.com/gin-gonic/gin"
)

func Router(db *users.NewUserRepo) *gin.Engine{
	router := gin.Default()

	u := handle.ReqUser{}

	u.ConnectorDb(db)

	router.POST("/create", u.Create)
	router.GET("/read/:id", u.Read)
	router.PUT("/update/:id", u.Update)
	router.DELETE("/delete/:id", u.Delete)
	router.GET("/getAll/:limit/:offset", u.GetAll)


	return router
}

