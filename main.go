package main

import (
	"bank/repository"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
const(
	host="localhost"
	port=5432
	user="anousone"
	password="smpidpmr"
	dbname="test"
)

func main(){
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	my_customer := repository.NewCustomerRepositoryDB(db)
	_ = my_customer

	customers, err := my_customer.GetAll()	
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}