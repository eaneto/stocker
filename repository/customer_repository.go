package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
)

var customerId uint = 0
var customerIdMutex sync.RWMutex = sync.RWMutex{}
var customersMutex sync.RWMutex = sync.RWMutex{}
var customersByID map[uint]domain.CustomerEntity = make(map[uint]domain.CustomerEntity)

type CustomerRepository interface {
	Save(customer domain.CustomerEntity) error
	FindAll() []domain.CustomerEntity
}

type CustomerRepositoryInMemory struct{}

func NewCustomerRepository() CustomerRepository {
	return CustomerRepositoryInMemory{}
}

func (CustomerRepositoryInMemory) Save(customer domain.CustomerEntity) error {
	customerIdMutex.Lock()
	customerId = customerId + 1
	customerIdMutex.Unlock()

	customer.ID = customerId
	customersMutex.Lock()
	customersByID[customer.ID] = customer
	customersMutex.Unlock()
	return nil
}

func (CustomerRepositoryInMemory) FindAll() []domain.CustomerEntity {
	customers := make([]domain.CustomerEntity, len(customersByID))
	i := 0
	for _, customer := range customersByID {
		customers[i] = customer
		i = i + 1
	}
	return customers
}
