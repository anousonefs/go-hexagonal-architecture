// this is port ປັກສຽບ

package repository

// i don't have customer table but i have person table in my database postgresql
// sorry about that
// step 1 create table structure
type Customer struct {
	Id  int    `db:"id"`
	First_name   string `db:"first_name"`
	Last_name string `db:"last_name"`
	Gender      string `db:"gender"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_customer_repository.go bank/repository CustomerRepository
// step 2 create role
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}