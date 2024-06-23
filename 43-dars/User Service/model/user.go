package model

type CreateUser struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Age    int    `json:"age"`
}
