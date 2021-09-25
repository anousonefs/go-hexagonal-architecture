package service

type NewAccountRequest struct {
	AccountType string `json:"account_type"`
	Amount int `json:"amount"`	
}

type AccountResponse struct {
	AccountID int   `json:"AccountID"`
	OpeningDate string   `json:"OpeningDate"`
	AccountType string   `json:"AccountType"`
	Amount int   `json:"Amount"`
	Status int   `json:"Status"`
}

type AccountService interface {
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts(int)([]AccountResponse, error)
}