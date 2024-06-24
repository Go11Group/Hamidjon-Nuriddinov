package model

type CreateUser struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Age    int    `json:"age"`
}

type UserBalance struct{
	UserId string `json:"user_id"`
	CardId string `json:"card_id"`
	CardNumber string `json:"number"`
	Balance float64 `json:"balance"`
}
