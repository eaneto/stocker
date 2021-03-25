package customer

import (
	"time"
)

type BaseCustomerService interface {
	Create(name string)
	FindAll() []Customer
}

type CustomerService struct {
	CustomerRepository CustomerRepository
}

func NewCustomerService() BaseCustomerService {
	return CustomerService{
		CustomerRepository: NewCustomerRepository(),
	}
}

func (service CustomerService) Create(name string) {
	customer := CustomerEntity{
		Name:      name,
		CreatedAt: time.Now(),
	}
	service.CustomerRepository.Save(customer)
}

func (service CustomerService) FindAll() []Customer {
	customerEntities := service.CustomerRepository.FindAll()
	customers := make([]Customer, len(customerEntities))
	for i, entity := range customerEntities {
		customers[i] = Customer{
			ID:   entity.ID,
			Name: entity.Name,
		}
	}
	return customers
}
