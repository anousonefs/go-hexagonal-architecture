package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository{
	customers := []Customer{
		{Id: 1001, First_name: "anousone", Last_name: "freestyle", Gender: "Male"},
		{Id: 1002, First_name: "sone", Last_name: "ff", Gender: "Male"},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll()([]Customer, error){
	return r.customers, nil
}

func(r customerRepositoryMock) GetById(id int)(*Customer, error){
	// should use binary search or binary search tree. The time complexcity is O(Log N)
	// But for loop Big-O is O(N)
	for _, customer := range r.customers {
		if customer.Id == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}