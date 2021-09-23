// this is adapter ຫົວປັກສຽບ

package repository

import (
	"github.com/jmoiron/sqlx"
)

// step 1 create struct with some field that you need to perform E.g db
// private
type customerRepositoryDB struct{
	db *sqlx.DB
}

// step 2 create new instance and return that port
func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository{
	// customerRepository is same type with CustomerRepository
	// because is implemented
	return customerRepositoryDB{ db: db }
}

// add method follow CustomerRepository Interface
func (r customerRepositoryDB) GetAll() ([]Customer, error){
	customers := []Customer{} // store data
	query := "SELECT id, first_name, last_name, gender FROM person LIMIT 10" // sql statement
	err := r.db.Select(&customers, query) // exec and assign to store data
	if err != nil {
		return nil, err
	}
	// return store data, error
	return customers, nil
}

// add method follow CustomerRepository Interface
func (r customerRepositoryDB) GetById(id int) (*Customer, error){
	customer := Customer{}
	query := "SELECT id, first_name, last_name, gender FROM person WHERE id=$1"
	err := r.db.Get(&customer,query, id)
	if err != nil{
		return nil, err
	}
	return &customer, nil
}




