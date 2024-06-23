package postgres

import (
	"database/sql"
	"mymode/model"
)

type TransactionRepo struct {
	Db *sql.DB
}


func NewTransactionRepo(db *sql.DB) *TransactionRepo{
	return &TransactionRepo{Db: db}
}


func (T *TransactionRepo) CreateTransaction(transaction model.CreateTransaction)error{
	_, err := T.Db.Exec(`INSERT INTO Transactions(amount, card_id, terminal_id, transaction_type) values($1, $2, $3, $4)`, transaction.Amount, transaction.CardId, transaction.TerminalId, transaction.TransactionType)
	return err
}


func (T *TransactionRepo) GetByIdTransaction(id string) (model.CreateTransaction, error){
	transaction := model.CreateTransaction{}
	err := T.Db.QueryRow(`SELECT * FROM Transactions WHERE transaction_id = $1`, id).Scan(&transaction.TransactionId, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
	return transaction, err
}


func (T *TransactionRepo) GetAllTransactions() ([]model.CreateTransaction, error){
	transactions := []model.CreateTransaction{}
	rows, err := T.Db.Query(`SELECT * FROM Transactions`)
	if err != nil{
		return transactions, err
	}
	for rows.Next(){
		transaction := model.CreateTransaction{}
		err := rows.Scan(&transaction.TransactionId, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
		if err != nil{
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err
}


func (T *TransactionRepo) UpdateTransaction(id string, transaction model.CreateTransaction) error{
	_, err := T.Db.Exec(`UPDATE Transactions SET card_id = $1, amount = $2, terminal_id = $3, transaction_type = $4 WHERE transaction_id = $5`, transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType, transaction.TransactionId)
	return err
}


func (T *TransactionRepo) DeleteTransaction(id string) error{
	_, err := T.Db.Exec(`DELETE FROM Transactions WHERE transaction_id = $1`, id)
	return err
}
