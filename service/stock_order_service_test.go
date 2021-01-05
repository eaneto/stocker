package service

import (
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StockOrderRepositoryMock struct {
	mock.Mock
}

func (m *StockOrderRepositoryMock) Save(stockOrder domain.StockOrderEntity) error {
	args := m.Called(stockOrder)
	return args.Error(0)
}

func (m *StockOrderRepositoryMock) Update(stockOrder domain.StockOrderEntity) error {
	args := m.Called(stockOrder)
	return args.Error(0)
}

func (m *StockOrderRepositoryMock) FindByCode(code uuid.UUID) (domain.StockOrderEntity, error) {
	args := m.Called(code)
	return args.Get(0).(domain.StockOrderEntity), args.Error(1)
}

func (m *StockOrderRepositoryMock) FindAllByCustomer(customerID uint) []domain.StockOrderEntity {
	args := m.Called(customerID)
	return args.Get(0).([]domain.StockOrderEntity)
}

type StockServiceMock struct {
	mock.Mock
}

func (m *StockServiceMock) RegisterStock(stock domain.Stock) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockServiceMock) SearchByTicker(ticker string) (domain.StockEntity, error) {
	args := m.Called(ticker)
	return args.Get(0).(domain.StockEntity), args.Error(1)
}
func (m *StockServiceMock) FindAll() []domain.Stock {
	args := m.Called()
	return args.Get(0).([]domain.Stock)
}

func (m *StockServiceMock) FindByID(id uint) (domain.StockEntity, error) {
	args := m.Called(id)
	return args.Get(0).(domain.StockEntity), args.Error(1)
}

func TestGetCustomerPositionWithNoneRegisteredStockShouldReturnEmptyPosition(t *testing.T) {
	stockOrderRepositoryMock := new(StockOrderRepositoryMock)
	stockServiceMock := new(StockServiceMock)

	orderService := StockOrderService{
		StockOrderRepository: stockOrderRepositoryMock,
		StockService:         stockServiceMock,
	}

	customerID := uint(1)
	stockOrderRepositoryMock.On("FindAllByCustomer", customerID).
		Return([]domain.StockOrderEntity{})

	_, err := orderService.GetCustomerPosition(customerID)

	assert.Nil(t, err)
}
