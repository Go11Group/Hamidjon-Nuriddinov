package packages

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Filter struct {
	Age           int
	FirstName     string
	LastName      string
	Nation        string
	Field         string
	Limit, Offset int
}

type NewRepoF struct {
	Db *sql.DB
}

func NewRepo(db *sql.DB) *NewRepoF {
	return &NewRepoF{Db: db}
}

func (F *NewRepoF) GetAll(c *gin.Context) {
	filter := Filter{}
	// params := make(map[string]interface{})
	// arr := []interface{}
	// limit := ""
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)

	query := `select first_name, last_name, age, nation, field from People where true`

	filt := ""

	filter.FirstName = c.Query("FirstName")
	filter.LastName = c.Query("LastName")
	filter.Field = c.Query("Field")
	filter.Nation = c.Query("Nation")
	if c.Query("Age") != ""{
		age, err := strconv.Atoi(c.Query("Age"))
		if err != nil{
			log.Fatal(err)
		}
		filter.Age = age
	}

	if len(filter.FirstName) > 0 {
		params["first_name"] = filter.FirstName
		filt += " and first_name = :first_name"
	}

	if len(filter.LastName) > 0 {
		params["last_name"] = filter.LastName
		filt += " and last_name = :last_name"
	}

	if len(filter.Field) > 0 {
		params["field"] = filter.Field
		filt += " and field = :field"
	}

	if len(filter.Nation) > 0 {
		params["nation"] = filter.Nation
		filt += " and nation = :nation"
	}

	if filter.Age > 0 {
		params["age"] = filter.Age
		filt += " and age = :age"
	}

	if filter.Limit > 0 {
		params["Limit"] = filter.Limit
		limit = ` Limit :limit`
	}

	
	query = query + filt + limit
	query, arr = ReplaceQueryParams(query, params)
	fmt.Println(query)
	rows, err := F.Db.Query(query, arr...)
	if err != nil {
		log.Fatal(err)
	}
	peoples := []Filter{}
	for rows.Next() {
		people := Filter{}
		rows.Scan(&people.FirstName, &people.LastName, &people.Nation, &people.Field, &people.Age)
		fmt.Println(people)
		peoples = append(peoples, people)
	}
	c.JSON(400, peoples)
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}
	return namedQuery, args
}
