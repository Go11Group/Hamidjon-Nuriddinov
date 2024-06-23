package model

type CreateCard struct {
	CardId string `json:"card_id"`
	Number string `json:"number"`
	UserId string `json:"user_id"`
}

type CreateStation struct {
	StationId string `json:"station_id"`
	Name      string `json:"name"`
}

type CreateTransaction struct {
	TransactionId   string  `json:"transaction_id"`
	CardId          string  `json:"card_id"`
	Amount          float64 `json:"amount"`
	TerminalId      string  `json:"terminal_id"`
	TransactionType string  `json:"transaction_type"`
}

type CreateTerminal struct {
	TerminalId string `json:"terminal_id"`
	StationId  string `json:"station_id"`
}

type UserBalance struct{
	UserId string `json:"user_id"`
	CardId string `json:"card_id"`
	CardNumber string `json:"number"`
	Balance float64 `json:"balance"`
}
