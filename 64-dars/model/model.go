package model

import "time"

type Company struct {
	Company string
	Price   int
	Time    time.Time
}

type RespCompany struct{
	Company string
	MinPrice int
	MaxPrice int
}

type CompanyName struct{
	Name string
}
