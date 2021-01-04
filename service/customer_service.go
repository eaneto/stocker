package service

import (
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
)

type BaseCustomerService interface {
	Create(name string)
	FindAll() []domain.Customer
}

type CustomerService struct {
	CustomerRepository repository.BaseCustomerRepository
}

func NewCustomerService() BaseCustomerService {
	return CustomerService{
		CustomerRepository: repository.NewCustomerRepository(),
	}
}

func (service CustomerService) Create(name string) {
	customer := domain.CustomerEntity{
		Name:      name,
		CreatedAt: time.Now(),
	}
	service.CustomerRepository.Save(customer)
}

func (service CustomerService) FindAll() []domain.Customer {
	customerEntities := service.CustomerRepository.FindAll()
	customers := make([]domain.Customer, len(customerEntities))
	for i, entity := range customerEntities {
		customers[i] = domain.Customer{
			ID:   entity.ID,
			Name: entity.Name,
		}
	}
	return customers
}
