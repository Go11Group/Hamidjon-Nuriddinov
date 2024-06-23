package model

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Age    int    `json:"age"`
}

type Card struct {
	CardId string `json:"card_id"`
	Number string `json:"number"`
	UserId string `json:"user_id"`
}

type Station struct {
	StationId string `json:"station_id"`
	Name      string `json:"name"`
}

type Transaction struct {
	TransactionId   string  `json:"transaction_id"`
	CardId          string  `json:"card_id"`
	Amount          float64 `json:"amount"`
	TerminalId      string  `json:"terminal_id"`
	TransactionType string  `json:"transaction_type"`
}

type Terminal struct {
	TerminalId string `json:"terminal_id"`
	StationId  string `json:"station_id"`
}
