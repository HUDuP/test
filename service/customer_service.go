package service

import (
	"bank/repository"
	"database/sql"
	"errors"
	"log"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{customerRepository: customerRepository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	 return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		
		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name: customer.Name,
		Status: customer.Status,
	}
	return &custResponse, nil
}
