package service

type CustomerResponse struct {
	Name string `json:"First_name"`
	Gender string `json:"Gender"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}