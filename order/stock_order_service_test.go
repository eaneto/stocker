package order

import (
	"testing"
	"time"

	"github.com/eaneto/stocker/stock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StockOrderRepositoryMock struct {
	mock.Mock
}

func (m *StockOrderRepositoryMock) Save(stockOrder StockOrderEntity) error {
	args := m.Called(stockOrder)
	return args.Error(0)
}

func (m *StockOrderRepositoryMock) Update(stockOrder StockOrderEntity) error {
	args := m.Called(stockOrder)
	return args.Error(0)
}

func (m *StockOrderRepositoryMock) FindByCode(code uuid.UUID) (StockOrderEntity, error) {
	args := m.Called(code)
	return args.Get(0).(StockOrderEntity), args.Error(1)
}

func (m *StockOrderRepositoryMock) FindAllByCustomer(customerID uint) []StockOrderEntity {
	args := m.Called(customerID)
	return args.Get(0).([]StockOrderEntity)
}

type StockServiceMock struct {
	mock.Mock
}

func (m *StockServiceMock) RegisterStock(stock stock.Stock) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockServiceMock) SearchByTicker(ticker string) (stock.StockEntity, error) {
	args := m.Called(ticker)
	return args.Get(0).(stock.StockEntity), args.Error(1)
}
func (m *StockServiceMock) FindAll() []stock.Stock {
	args := m.Called()
	return args.Get(0).([]stock.Stock)
}

func (m *StockServiceMock) FindByID(id uint) (stock.StockEntity, error) {
	args := m.Called(id)
	return args.Get(0).(stock.StockEntity), args.Error(1)
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
		Return([]StockOrderEntity{})

	position, err := orderService.GetCustomerPosition(customerID)

	assert.Nil(t, err)
	assert.Empty(t, position.Stocks)
}

func TestGetCustomerPositionWithOneRegisteredStockShouldReturnPositionWithThisStock(t *testing.T) {
	stockOrderRepositoryMock := new(StockOrderRepositoryMock)
	stockServiceMock := new(StockServiceMock)

	orderService := StockOrderService{
		StockOrderRepository: stockOrderRepositoryMock,
		StockService:         stockServiceMock,
	}
	customerID := uint(1)
	code, _ := uuid.NewRandom()
	orders := []StockOrderEntity{
		{
			Code:       code,
			StockID:    1,
			CustomerID: customerID,
			Amount:     100,
			Status:     Created,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	stock := stock.StockEntity{
		ID:     orders[0].StockID,
		Ticker: "STCK9",
		Price:  1000,
	}

	stockOrderRepositoryMock.On("FindAllByCustomer", customerID).
		Return(orders)
	stockServiceMock.On("FindByID", orders[0].StockID).Return(stock, nil)

	position, err := orderService.GetCustomerPosition(customerID)

	assert.Nil(t, err)
	assert.Equal(t, position.CustomerID, customerID)
	assert.NotEmpty(t, position.Stocks)
	assert.Equal(t, len(orders), len(position.Stocks))
	stockPosition := position.Stocks[0]
	assert.Equal(t, stock.Price, stockPosition.Price)
	assert.Equal(t, stock.Ticker, stockPosition.Ticker)
	assert.Equal(t, orders[0].Amount, stockPosition.Amount)
}
