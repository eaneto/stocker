package service

import (
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) Save(stock domain.CustomerEntity) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *CustomerRepositoryMock) FindByCode(code uuid.UUID) (domain.CustomerEntity, error) {
	args := m.Called(code)
	return args.Get(0).(domain.CustomerEntity), args.Error(1)
}

func (m *CustomerRepositoryMock) FindAll() []domain.CustomerEntity {
	args := m.Called()
	return args.Get(0).([]domain.CustomerEntity)
}

func TestRegisterCustomerWithSuccessShouldCallRepository(t *testing.T) {
	repository := new(CustomerRepositoryMock)

	name := "Edison"
	repository.On("Save", mock.Anything).Return(nil)

	service := CustomerService{
		CustomerRepository: repository,
	}

	service.Create(name)

	repository.AssertExpectations(t)
}

func TestFindAllCustomersReturningEmptyShouldReturnEmptyList(t *testing.T) {
	repository := new(CustomerRepositoryMock)

	customers := []domain.CustomerEntity{}
	repository.On("FindAll").Return(customers)

	service := CustomerService{
		CustomerRepository: repository,
	}

	foundCustomers := service.FindAll()

	assert.Empty(t, foundCustomers)
}

func TestFindAllCustomersReturningOneItemShouldReturnListWithOneItem(t *testing.T) {
	repository := new(CustomerRepositoryMock)

	customers := []domain.CustomerEntity{
		{
			Name: "ABV9",
		},
	}
	repository.On("FindAll").Return(customers)

	service := CustomerService{
		CustomerRepository: repository,
	}

	foundCustomers := service.FindAll()

	assert.Equal(t, len(customers), len(foundCustomers))
	assert.Equal(t, customers[0].Name, foundCustomers[0].Name)
}
