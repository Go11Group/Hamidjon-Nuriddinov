package postgres

import (
	"database/sql"
	"mymode/model"
)

type TerminalRepo struct {
	Db *sql.DB
}


func NewTerminalRepo(db *sql.DB) *TerminalRepo{
	return &TerminalRepo{Db: db}
}


func (T *TerminalRepo) CreateTerminal(terminal model.CreateTerminal)error{
	_, err := T.Db.Exec(`INSERT INTO Terminals(station_id) values($1)`, terminal.StationId)
	return err
}


func (T *TerminalRepo) GetByIdTerminal(id string) (model.CreateTerminal, error){
	terminal := model.CreateTerminal{}
	err := T.Db.QueryRow(`SELECT * FROM Terminals WHERE terminal_id = $1`, id).Scan(&terminal.TerminalId, &terminal.StationId)
	return terminal, err
}


func (T *TerminalRepo) GetAllTerminals() ([]model.CreateTerminal, error){
	terminals := []model.CreateTerminal{}
	rows, err := T.Db.Query(`SELECT * FROM Terminals`)
	if err != nil{
		return terminals, err
	}
	for rows.Next(){
		terminal := model.CreateTerminal{}
		err := rows.Scan(&terminal.TerminalId, &terminal.StationId)
		if err != nil{
			return terminals, err
		}
		terminals = append(terminals, terminal)
	}
	return terminals, err
}


func (T *TerminalRepo) UpdateTerminal(id string, terminal model.CreateTerminal) error{
	_, err := T.Db.Exec(`UPDATE Terminals SET station_id = $1 WHERE terminal_id = $2`, terminal.StationId, terminal.TerminalId)
	return err
}


func (T *TerminalRepo) DeleteTerminal(id string) error{
	_, err := T.Db.Exec(`DELETE FROM Terminals WHERE terminal_id = $1`, id)
	return err
}
