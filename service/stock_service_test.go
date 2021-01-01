package service

import (
	"errors"
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StockRepositoryMock struct {
	mock.Mock
}

func (m *StockRepositoryMock) Save(stock domain.StockEntity) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockRepositoryMock) FindByTicker(ticker string) (domain.StockEntity, error) {
	args := m.Called(ticker)
	return args.Get(0).(domain.StockEntity), args.Error(1)
}

func (m *StockRepositoryMock) FindAll() []domain.StockEntity {
	args := m.Called()
	return args.Get(0).([]domain.StockEntity)
}

func TestRegisterStockWithSuccessShouldReturnNilError(t *testing.T) {
	repository := new(StockRepositoryMock)

	stock := domain.Stock{
		Ticker: "ABV9",
		Price:  100,
	}

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

	stock := domain.Stock{
		Ticker: "ABV9",
		Price:  100,
	}

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
	stockEntity := domain.StockEntity{
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
	stockEntity := domain.StockEntity{}

	notFoundError := domain.StockNotFoundError{}
	repository.On("FindByTicker", ticker).Return(stockEntity, notFoundError)

	service := StockService{
		StockRepository: repository,
	}

	_, err := service.SearchByTicker(ticker)

	assert.NotNil(t, err)
	_, isNotFound := err.(domain.StockNotFoundError)
	assert.True(t, isNotFound)
	repository.AssertExpectations(t)
}

func TestFindAllReturningEmptyShouldReturnEmptyList(t *testing.T) {
	repository := new(StockRepositoryMock)

	stocks := []domain.StockEntity{}
	repository.On("FindAll").Return(stocks)

	service := StockService{
		StockRepository: repository,
	}

	foundStocks := service.FindAll()

	assert.Empty(t, foundStocks)
}

func TestFindAllReturningOneItemShouldReturnListWithOneItem(t *testing.T) {
	repository := new(StockRepositoryMock)

	stocks := []domain.StockEntity{
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
