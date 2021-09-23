package repository

type Account struct {
	AccountID int   `db:"AccountID"`
	CustomerID int   `db:"CustomerID"`
	OpeningDate string   `db:"OpeningDate"`
	AccountType string   `db:"AccountType"`
	Amount float64   `db:"Amount"`
	Status int   `db:"Status"`
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}