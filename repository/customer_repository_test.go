package repository

import (
	"testing"
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveCustomerShouldSaveOnMap(t *testing.T) {
	clearAllCustomers()

	repository := CustomerRepository{}

	customer := domain.CustomerEntity{
		Name:      "Edison",
		CreatedAt: time.Now(),
	}

	err := repository.Save(customer)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(customersByID))
	foundCustomer := customersByID[customerId]
	assert.Equal(t, customer.Name, foundCustomer.Name)
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

	customers := make(map[uint]domain.CustomerEntity)
	id := uint(1)
	customer := domain.CustomerEntity{
		Name: "Françoise",
		ID:   id,
	}
	customers[id] = customer
	customersByID = customers

	repository := CustomerRepository{}

	foundCustomers := repository.FindAll()

	assert.Equal(t, len(customers), len(foundCustomers))
	assert.Equal(t, customer, foundCustomers[0])
}

// clearAllCustomers Clears all stores customers and resets ids.
func clearAllCustomers() {
	customerId = 0
	customersByID = make(map[uint]domain.CustomerEntity)
}
