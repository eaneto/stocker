package controller

import (
	"net/http"
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerServiceMock struct {
	mock.Mock
}

func (m *CustomerServiceMock) Create(name string) {
	m.Called(name)
}

func (m *CustomerServiceMock) FindByCode(code uuid.UUID) (domain.CustomerEntity, error) {
	args := m.Called(code)
	return args.Get(0).(domain.CustomerEntity), args.Error(1)
}

func (m *CustomerServiceMock) FindAll() []domain.Customer {
	args := m.Called()
	return args.Get(0).([]domain.Customer)
}

func TestCreateCustomerShouldReturnCreated(t *testing.T) {
	service := new(CustomerServiceMock)

	name := "Edison"
	service.On("Create", name)

	controller := CustomerController{
		CustomerService: service,
	}

	status := controller.Create(name)

	assert.Equal(t, http.StatusCreated, status)
}

func TestFindAllCustomersReturnListWithCustomersAndOk(t *testing.T) {
	service := new(CustomerServiceMock)

	service.On("FindAll").Return([]domain.Customer{})

	controller := CustomerController{
		CustomerService: service,
	}

	customers, status := controller.FindAll()

	assert.Empty(t, customers)
	assert.Equal(t, http.StatusOK, status)
}
