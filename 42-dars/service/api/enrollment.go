package api

import (
	"log"
	"mymode/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReqEnrollment struct {
	Db Handle
}

func (E *ReqEnrollment) ConnectorDb(db Handle) {
	E.Db = db
}

func (E *ReqEnrollment) CreateEnrollment(c *gin.Context) {
	enrollment := model.Enrollment{}
	err := c.ShouldBindJSON(&enrollment)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, err)
		return
	}
	err = E.Db.e.CreateEnrollment(enrollment)
	if err != nil {
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
	enrollment := model.Enrollment{}

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
	params := make(map[string]interface{})
	filter := ""
	enrollmentFilter := model.EnrollmentFilter{}
	var err error
	enrollmentFilter.UserId = c.Query("user_id")
	enrollmentFilter.CourseId = c.Query("course_id")
	enrollmentFilter.EnrollmentDate = c.Query("enrollment_date")
	limit := c.Query("limit")
	if limit != "" {
		enrollmentFilter.Limit, err = strconv.Atoi(limit)
		if err != nil {
			enrollmentFilter.Limit = 0
		}
	} else {
		enrollmentFilter.Limit = 0
	}
	offset := c.Query("offset")
	if offset != "" {
		enrollmentFilter.Offset, err = strconv.Atoi(offset)
		if err != nil {
			enrollmentFilter.Offset = 0
		}
	} else {
		enrollmentFilter.Offset = 0
	}

	if len(enrollmentFilter.UserId) > 0 {
		params["user_id"] = enrollmentFilter.UserId
		filter += " and user_id = :user_id"
	}

	if len(enrollmentFilter.CourseId) > 0 {
		params["course_id"] = enrollmentFilter.CourseId
		filter += " and course_id = :course_id"
	}

	if len(enrollmentFilter.EnrollmentDate) > 0 {
		params["enrollment_date"] = enrollmentFilter.EnrollmentDate
		filter += " and enrollment_date = :enrollment_date"
	}

	if enrollmentFilter.Limit > 0 {
		params["limit"] = enrollmentFilter.Limit
		filter += " limit = :limit"
	}

	if enrollmentFilter.Offset > 0 {
		params["offset"] = enrollmentFilter.Offset
		filter += " offset = :offset"
	}

	query, arr := ReplaceQueryParamsEnrollment(filter, params)

	enrollments, err := E.Db.e.GetAllEnrollments(query, arr)
	if err != nil {
		log.Fatal(err)
		c.JSON(404, err)
		return
	}
	c.JSON(200, enrollments)
}

func ReplaceQueryParamsEnrollment(query string, params map[string]interface{}) (string, []interface{}) {
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
