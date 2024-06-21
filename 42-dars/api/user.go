package api

import (
	"log"
	"mymode/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReqUser struct {
	Db Handle
}

func (U *ReqUser) ConnectorDb(db Handle) {
	U.Db = db
}

func (U *ReqUser) CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}

	
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
	user := model.User{}

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
	params := make(map[string]interface{})
	filter := ""
	userFilter := model.UserFilter{}
	var err error
	userFilter.Name = c.Query("name")
	userFilter.Birthday = c.Query("birthday")
	userFilter.Email = c.Query("email")
	limit := c.Query("limit")
	if limit != "" {
		userFilter.Limit, err = strconv.Atoi(limit)
		if err != nil {
			userFilter.Limit = 0
		}
	} else {
		userFilter.Limit = 0
	}
	offset := c.Query("offset")
	if offset != "" {
		userFilter.Offset, err = strconv.Atoi(offset)
		if err != nil {
			userFilter.Offset = 0
		}
	} else {
		userFilter.Offset = 0
	}

	if len(userFilter.Name) > 0 {
		params["name"] = userFilter.Name
		filter += " and name = :name"
	}

	if len(userFilter.Birthday) > 0 {
		params["birthday"] = userFilter.Birthday
		filter += " and birthday = :birthday"
	}

	if len(userFilter.Email) > 0 {
		params["email"] = userFilter.Email
		filter += " and email = :email"
	}

	if userFilter.Limit > 0 {
		params["limit"] = userFilter.Limit
		filter += " limit = :limit"
	}

	if userFilter.Offset > 0 {
		params["offset"] = userFilter.Offset
		filter += " offset = :offset"
	}

	query, arr := ReplaceQueryParamsUser(filter, params)

	users, err := U.Db.u.GetAllUsers(query, arr)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, users)
}

func ReplaceQueryParamsUser(query string, params map[string]interface{}) (string, []interface{}) {
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

func (U *ReqUser) GetCoursesByUser(c *gin.Context) {
	id := c.Param("user_id")
	userCourses, err := U.Db.u.GetCoursesByUser(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, userCourses)
}

func (U *ReqUser) SearchUsers(c *gin.Context) {
	params := make(map[string]interface{})
	filter := ""
	var err error

	user := model.UserSearch{}
	user.Name = c.Query("name")
	user.Email = c.Query("email")

	if c.Query("age_from") != "" {
		user.AgeFrom, err = strconv.Atoi(c.Query("age_from"))
		if err != nil {
			user.AgeFrom = 0
		}
	} else {
		user.AgeFrom = 0
	}
	if c.Query("age_to") != "" {
		user.AgeTo, err = strconv.Atoi(c.Query("age_to"))
		if err != nil {
			user.AgeTo = 150
		}
	} else {
		user.AgeTo = 150
	}

	if len(user.Name) > 0 {
		params["name"] = user.Name
		filter += " and name = :name"
	}
	if len(user.Email) > 0 {
		params["email"] = user.Email
		filter += " and email = :email"
	}
		params["age_from"] = user.AgeFrom
		params["age_to"] = user.AgeTo

	filter += " AND EXTRACT(YEAR FROM age(birthday)) BETWEEN :age_from AND :age_to"
	filter, arr := ReplaceQueryParamsUser(filter, params)

	users, err := U.Db.u.SearchUsers(filter, arr)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, users)
}

