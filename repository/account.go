package repository

type Account struct {
	AccountID int   `db:"accountid"`
	CustomerID int   `db:"customerid"`
	OpeningDate string   `db:"openingdate"`
	AccountType string   `db:"accounttype"`
	Amount int   `db:"amount"`
	Status int   `db:"status"`
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}