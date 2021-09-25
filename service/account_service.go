package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService{
	return accountService{accRepo: accRepo}
}

func (a accountService) NewAccount(customerID int, accRequest NewAccountRequest) (*AccountResponse, error){
	//* Validate
	if accRequest.Amount < 5000 {
		return nil, errs.NewValidateError("amount at least 5,000")
	}
	if strings.ToLower(accRequest.AccountType) != "saving" && strings.ToLower(accRequest.AccountType) != "checking"{
		return nil, errs.NewValidateError("account type should by saving or checking")
	}

	account := repository.Account{
		AccountID: 0,
		CustomerID: customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: accRequest.AccountType,
		Amount: accRequest.Amount,
		Status: 1,
	}
	// data from repository
	newAcc, err := a.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	// response data
	accResponse := AccountResponse{
		AccountID: newAcc.AccountID,
		AccountType: newAcc.AccountType,
		Amount: newAcc.Amount,
		OpeningDate: newAcc.OpeningDate,
		Status: newAcc.Status,
	}
	return &accResponse, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error){
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	accResponses := []AccountResponse{}

	for _, v := range accounts{
		accResponse := AccountResponse{
			AccountID: v.AccountID,
			OpeningDate: v.OpeningDate ,
			AccountType: v.AccountType,
			Amount: v.Amount,
			Status: v.Status,
		}
		accResponses = append(accResponses, accResponse)
	}

	return accResponses, nil
}