package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository{
	return accountRepositoryDB{db: db}
}

func (a accountRepositoryDB) Create(acc Account) (*Account, error){
	stmt := "INSERT INTO account (CustomerID, OpeningDate, AccountType, Amount, Status) VALUES ($1, $2, $3, $4, $5)"	
	result, err := a.db.Exec(
		stmt,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()	
	if err != nil {
		fmt.Print("lastinsert id failed")
		fmt.Println(id)
		return nil, err
	}
	acc.AccountID = int(id)
	return &acc, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error){
	query := "SELECT AccountID, CustomerID, OpeningDate, AccountType, Amount, Status FROM account WHERE CustomerID=$1"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, nil

}