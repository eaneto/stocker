package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
)

var customerId uint = 0
var customerIdMutex sync.Mutex = sync.Mutex{}
var customersMutex sync.Mutex = sync.Mutex{}
var customersByCode map[uuid.UUID]domain.CustomerEntity = make(map[uuid.UUID]domain.CustomerEntity)

type BaseCustomerRepository interface {
	Save(customer domain.CustomerEntity) error
	FindAll() []domain.CustomerEntity
}

type CustomerRepository struct{}

func NewCustomerRepository() BaseCustomerRepository {
	return CustomerRepository{}
}

func (CustomerRepository) Save(customer domain.CustomerEntity) error {
	customerIdMutex.Lock()
	customerId = customerId + 1
	customerIdMutex.Unlock()

	customer.ID = customerId
	customersMutex.Lock()
	customersByCode[customer.Code] = customer
	customersMutex.Unlock()
	return nil
}

func (CustomerRepository) FindAll() []domain.CustomerEntity {
	customers := make([]domain.CustomerEntity, len(customersByCode))
	i := 0
	for _, customer := range customersByCode {
		customers[i] = customer
		i = i + 1
	}
	return customers
}
