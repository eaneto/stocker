package customer

import (
	"sync"
)

var customerId uint = 0
var customerIdMutex sync.RWMutex = sync.RWMutex{}
var customersMutex sync.RWMutex = sync.RWMutex{}
var customersByID map[uint]CustomerEntity = make(map[uint]CustomerEntity)

type CustomerRepository interface {
	Save(customer CustomerEntity) error
	FindAll() []CustomerEntity
}

type CustomerRepositoryInMemory struct{}

func NewCustomerRepository() CustomerRepository {
	return CustomerRepositoryInMemory{}
}

func (CustomerRepositoryInMemory) Save(customer CustomerEntity) error {
	customerIdMutex.Lock()
	customerId++
	customerIdMutex.Unlock()

	customer.ID = customerId
	customersMutex.Lock()
	customersByID[customer.ID] = customer
	customersMutex.Unlock()
	return nil
}

func (CustomerRepositoryInMemory) FindAll() []CustomerEntity {
	customers := make([]CustomerEntity, len(customersByID))
	i := 0
	for _, customer := range customersByID {
		customers[i] = customer
		i++
	}
	return customers
}
