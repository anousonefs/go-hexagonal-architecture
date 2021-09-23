package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	// type interface from repository
	custRepo repository.CustomerRepository
}

// recieve instance interface of repository
func NewCustomerService(custRepo repository.CustomerRepository) CustomerService{
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error){
	// from back
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// business
	custResponses := []CustomerResponse{}
	for _, v := range customers {
		custResponse := CustomerResponse{
			Name: v.First_name,
			Gender: v.Gender,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error){
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found Der!!")
		}
		// print on terminal
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	cusResponse := CustomerResponse{
		Name: customer.First_name,
		Gender: customer.Gender,
	}
	return &cusResponse, nil
}