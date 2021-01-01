package usecase

import (
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StockServiceMock struct {
	mock.Mock
}

func (m *StockServiceMock) RegisterStock(stock domain.Stock) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockServiceMock) SearchByTicker(ticker string) (domain.Stock, error) {
	args := m.Called(ticker)
	return args.Get(0).(domain.Stock), args.Error(1)
}

func TestRegisterStockWithUnregisteredTickerShouldNotReturnError(t *testing.T) {
	service := new(StockServiceMock)

	stock := domain.Stock{
		Ticker: "ABC2",
		Price:  12300,
	}

	notFoundError := domain.StockNotFoundError{}
	service.On("SearchByTicker", stock.Ticker).Return(domain.Stock{}, notFoundError)
	service.On("RegisterStock", stock).Return(nil)

	usecase := StockUseCase{
		StockService: service,
	}

	err := usecase.RegisterStock(stock)

	assert.Nil(t, err)
	service.AssertExpectations(t)
}

func TestRegisterStockWithAlreadyRegisteredTickerShouldReturnError(t *testing.T) {
	service := new(StockServiceMock)

	stock := domain.Stock{
		Ticker: "ABC2",
		Price:  12300,
	}

	service.On("SearchByTicker", stock.Ticker).Return(stock, nil)

	usecase := StockUseCase{
		StockService: service,
	}

	err := usecase.RegisterStock(stock)

	assert.NotNil(t, err)
	_, isAlreadyRegistered := err.(domain.AlreadyRegisteredStockError)
	assert.True(t, isAlreadyRegistered)
	service.AssertExpectations(t)
	service.AssertNotCalled(t, "RegisterStock", mock.Anything)
}

func TestSearchStockByTickerFoundStockShouldBeReturned(t *testing.T) {
	service := new(StockServiceMock)

	stock := domain.Stock{
		Ticker: "ABC2",
		Price:  12300,
	}

	service.On("SearchByTicker", stock.Ticker).Return(stock, nil)

	usecase := StockUseCase{
		StockService: service,
	}

	foundStock, err := usecase.SearchByTicker(stock.Ticker)

	assert.Nil(t, err)
	service.AssertExpectations(t)
	assert.Equal(t, stock.Ticker, foundStock.Ticker)
	assert.Equal(t, stock.Price, foundStock.Price)
}

func TestSearchUnregisteredStockByTickerShouldReturnNotFoundError(t *testing.T) {
	service := new(StockServiceMock)

	stock := domain.Stock{
		Ticker: "ABC2",
		Price:  12300,
	}

	notFoundError := domain.StockNotFoundError{}
	service.On("SearchByTicker", stock.Ticker).Return(domain.Stock{}, notFoundError)

	usecase := StockUseCase{
		StockService: service,
	}

	foundStock, err := usecase.SearchByTicker(stock.Ticker)

	assert.NotNil(t, err)
	assert.Equal(t, notFoundError, err)
	service.AssertExpectations(t)
	assert.Empty(t, foundStock.Ticker)
	assert.Equal(t, uint(0), foundStock.Price)
}
