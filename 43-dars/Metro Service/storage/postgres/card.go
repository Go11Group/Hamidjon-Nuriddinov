package postgres

import (
	"database/sql"
	"mymode/model"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (C *CardRepo) CreateCard(card model.CreateCard) error {
	_, err := C.Db.Exec(`INSERT INTO Cards(number, user_id) values($1, $2)`, card.Number, card.UserId)
	return err
}

func (C *CardRepo) GetByIdCard(id string) (model.CreateCard, error) {
	card := model.CreateCard{}
	err := C.Db.QueryRow(`SELECT * FROM Cards WHERE card_id = $1`, id).Scan(&card.CardId, &card.Number, &card.UserId)
	return card, err
}

func (C *CardRepo) GetAllCards() ([]model.CreateCard, error) {
	cards := []model.CreateCard{}
	rows, err := C.Db.Query(`SELECT * FROM Cards`)
	if err != nil {
		return cards, err
	}
	for rows.Next() {
		card := model.CreateCard{}
		err := rows.Scan(&card.CardId, &card.Number, &card.UserId)
		if err != nil {
			return cards, err
		}
		cards = append(cards, card)
	}
	return cards, err
}

func (C *CardRepo) UpdateCard(id string, card model.CreateCard) error {
	_, err := C.Db.Exec(`UPDATE Cards SET number = $1, user_id = $2 WHERE card_id = $3`, card.Number, card.UserId, card.CardId)
	return err
}

func (C *CardRepo) DeleteCard(id string) error {
	_, err := C.Db.Exec(`DELETE FROM Cards WHERE card_id = $1`, id)
	return err
}


