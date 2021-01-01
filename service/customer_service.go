package service

import (
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
	"github.com/google/uuid"
)

type BaseCustomerService interface {
	Create(name string)
	FindByCode(code uuid.UUID) (domain.CustomerEntity, error)
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
	code, _ := uuid.NewRandom()
	customer := domain.CustomerEntity{
		Name:      name,
		Code:      code,
		CreatedAt: time.Now(),
	}
	service.CustomerRepository.Save(customer)
}

func (service CustomerService) FindByCode(code uuid.UUID) (domain.CustomerEntity, error) {
	return service.CustomerRepository.FindByCode(code)
}

func (service CustomerService) FindAll() []domain.Customer {
	customerEntities := service.CustomerRepository.FindAll()
	customers := make([]domain.Customer, len(customerEntities))
	for i, entity := range customerEntities {
		customers[i] = domain.Customer{
			Name: entity.Name,
			Code: entity.Code,
		}
	}
	return customers
}
