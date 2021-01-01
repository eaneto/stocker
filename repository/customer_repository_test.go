package repository

import (
	"testing"
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveCustomerShouldSaveOnMap(t *testing.T) {
	clearAllCustomers()

	repository := CustomerRepository{}

	code, _ := uuid.NewRandom()
	customer := domain.CustomerEntity{
		Name:      "Edison",
		Code:      code,
		CreatedAt: time.Now(),
	}

	err := repository.Save(customer)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(customersByCode))
	foundCustomer := customersByCode[customer.Code]
	assert.Equal(t, customer.Name, foundCustomer.Name)
	assert.Equal(t, customer.Code, foundCustomer.Code)
	assert.Equal(t, customer.CreatedAt, foundCustomer.CreatedAt)
}

func TestFindAllCustomersWithNoneRegisteredShouldReturnEmptySlice(t *testing.T) {
	clearAllCustomers()

	repository := CustomerRepository{}

	customers := repository.FindAll()

	assert.Empty(t, customers)
}

func TestFindAllCustomersWithOneRegisteredShouldReturnSliceWithOneElement(t *testing.T) {
	clearAllCustomers()

	customers := make(map[uuid.UUID]domain.CustomerEntity)
	code, _ := uuid.NewRandom()
	customer := domain.CustomerEntity{
		Name: "Françoise",
		Code: code,
	}
	customers[code] = customer
	customersByCode = customers

	repository := CustomerRepository{}

	foundCustomers := repository.FindAll()

	assert.Equal(t, len(customers), len(foundCustomers))
	assert.Equal(t, customer, foundCustomers[0])
}

func TestFindCustomerByCodeOfExistentCustomerShouldReturnCustomer(t *testing.T) {
	clearAllCustomers()

	customers := make(map[uuid.UUID]domain.CustomerEntity)
	code, _ := uuid.NewRandom()
	customer := domain.CustomerEntity{
		Name: "Françoise",
		Code: code,
	}
	customers[code] = customer
	customersByCode = customers

	repository := CustomerRepository{}

	foundCustomer, err := repository.FindByCode(code)

	assert.Nil(t, err)
	assert.Equal(t, customer, foundCustomer)
}

func TestFindCustomerByCodeOfNonExistentCustomerShouldReturnError(t *testing.T) {
	clearAllCustomers()

	code, _ := uuid.NewRandom()
	repository := CustomerRepository{}

	_, err := repository.FindByCode(code)

	assert.NotNil(t, err)
	_, isNotFound := err.(domain.CustomerNotFoundError)
	assert.True(t, isNotFound)
}

// clearAllCustomers Clears all stores customers and resets ids.
func clearAllCustomers() {
	customerId = 0
	customersByCode = make(map[uuid.UUID]domain.CustomerEntity)
}
