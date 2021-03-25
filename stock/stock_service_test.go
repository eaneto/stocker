package stock

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StockRepositoryMock struct {
	mock.Mock
}

func (m *StockRepositoryMock) Save(stock StockEntity) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockRepositoryMock) FindByTicker(ticker string) (StockEntity, error) {
	args := m.Called(ticker)
	return args.Get(0).(StockEntity), args.Error(1)
}

func (m *StockRepositoryMock) FindByID(id uint) (StockEntity, error) {
	args := m.Called(id)
	return args.Get(0).(StockEntity), args.Error(1)
}

func (m *StockRepositoryMock) FindAll() []StockEntity {
	args := m.Called()
	return args.Get(0).([]StockEntity)
}

func TestRegisterStockWithSuccessShouldReturnNilError(t *testing.T) {
	repository := new(StockRepositoryMock)

	stock := Stock{
		Ticker: "ABV9",
		Price:  100,
	}
	stockEntity := StockEntity{}

	repository.On("FindByTicker", stock.Ticker).
		Return(stockEntity, StockNotFoundError{})
	repository.On("Save", mock.Anything).Return(nil)

	service := StockService{
		StockRepository: repository,
	}

	err := service.RegisterStock(stock)

	assert.Nil(t, err)
	repository.AssertExpectations(t)
}

func TestRegisterStockWithErrorShouldReturnError(t *testing.T) {
	repository := new(StockRepositoryMock)

	stock := Stock{
		Ticker: "ABV9",
		Price:  100,
	}

	stockEntity := StockEntity{}

	repository.On("FindByTicker", stock.Ticker).
		Return(stockEntity, StockNotFoundError{})
	expectedError := errors.New("Error creating stock")
	repository.On("Save", mock.Anything).Return(expectedError)

	service := StockService{
		StockRepository: repository,
	}

	err := service.RegisterStock(stock)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	repository.AssertExpectations(t)
}

func TestSearchRegisteredStockByTickerShouldReturnStock(t *testing.T) {
	repository := new(StockRepositoryMock)

	ticker := "ABV9"
	stockEntity := StockEntity{
		ID:     uint(1),
		Ticker: ticker,
		Price:  100,
	}

	repository.On("FindByTicker", ticker).Return(stockEntity, nil)

	service := StockService{
		StockRepository: repository,
	}

	foundStock, err := service.SearchByTicker(ticker)

	assert.Nil(t, err)
	repository.AssertExpectations(t)
	assert.Equal(t, stockEntity.Ticker, foundStock.Ticker)
	assert.Equal(t, stockEntity.Price, foundStock.Price)
}

func TestSearchUneegisteredStockByTickerShouldReturnError(t *testing.T) {
	repository := new(StockRepositoryMock)

	ticker := "ABV9"
	stockEntity := StockEntity{}

	notFoundError := StockNotFoundError{}
	repository.On("FindByTicker", ticker).Return(stockEntity, notFoundError)

	service := StockService{
		StockRepository: repository,
	}

	_, err := service.SearchByTicker(ticker)

	assert.NotNil(t, err)
	_, isNotFound := err.(StockNotFoundError)
	assert.True(t, isNotFound)
	repository.AssertExpectations(t)
}

func TestFindAllReturningEmptyShouldReturnEmptyList(t *testing.T) {
	repository := new(StockRepositoryMock)

	stocks := []StockEntity{}
	repository.On("FindAll").Return(stocks)

	service := StockService{
		StockRepository: repository,
	}

	foundStocks := service.FindAll()

	assert.Empty(t, foundStocks)
}

func TestFindAllReturningOneItemShouldReturnListWithOneItem(t *testing.T) {
	repository := new(StockRepositoryMock)

	stocks := []StockEntity{
		{
			Ticker: "ABV9",
		},
	}
	repository.On("FindAll").Return(stocks)

	service := StockService{
		StockRepository: repository,
	}

	foundStocks := service.FindAll()

	assert.Equal(t, len(stocks), len(foundStocks))
	assert.Equal(t, stocks[0].Ticker, foundStocks[0].Ticker)
}
