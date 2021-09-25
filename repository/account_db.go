package repository

import (
	"github.com/jmoiron/sqlx"
)

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository{
	return accountRepositoryDB{db: db}
}

func (a accountRepositoryDB) Create(acc Account) (*Account, error){
	stmt := "INSERT INTO account (CustomerID, OpeningDate, AccountType, Amount, Status) VALUES ($1, $2, $3, $4, $5)RETURNING accountid"	
	LastInsertId := 0
	err := a.db.QueryRow(
		stmt,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	).Scan(&LastInsertId)
	if err != nil {
		return nil, err
	}
	acc.AccountID = int(LastInsertId)
	return &acc, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error){
	query := "SELECT accountid, customerid, openingdate, accounttype, amount, status FROM account WHERE customerid=$1"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, nil

}