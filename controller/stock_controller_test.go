package controller

import (
	"errors"
	"net/http"
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StockUseCaseMock struct {
	mock.Mock
}

func (m *StockUseCaseMock) RegisterStock(stock domain.Stock) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockUseCaseMock) SearchByTicker(ticker string) (domain.Stock, error) {
	args := m.Called(ticker)
	return args.Get(0).(domain.Stock), args.Error(1)
}

func TestRegisterUnregisteredStockShouldReturnCreated(t *testing.T) {
	stockUsecase := new(StockUseCaseMock)

	stock := domain.Stock{
		Ticker: "UYA2",
		Price:  100,
	}

	stockUsecase.On("RegisterStock", stock).Return(nil)

	controller := StockController{
		StockUseCase: stockUsecase,
	}

	status := controller.RegisterStock(stock)

	assert.Equal(t, status, http.StatusCreated)
}

func TestRegisterAlreadyRegisteredStockShouldReturnConflict(t *testing.T) {
	stockUsecase := new(StockUseCaseMock)

	stock := domain.Stock{
		Ticker: "UYA2",
		Price:  100,
	}

	error := domain.AlreadyRegisteredStockError{Ticker: stock.Ticker}
	stockUsecase.On("RegisterStock", stock).Return(error)

	controller := StockController{
		StockUseCase: stockUsecase,
	}

	status := controller.RegisterStock(stock)

	assert.Equal(t, status, http.StatusConflict)
}

func TestRegisterWithUnexpectedErrorShouldReturnInternalServerError(t *testing.T) {
	stockUsecase := new(StockUseCaseMock)

	stock := domain.Stock{
		Ticker: "UYA2",
		Price:  100,
	}

	error := errors.New("Unexpected error")
	stockUsecase.On("RegisterStock", stock).Return(error)

	controller := StockController{
		StockUseCase: stockUsecase,
	}

	status := controller.RegisterStock(stock)

	assert.Equal(t, status, http.StatusInternalServerError)
}

func TestSearchRegisteredStockShouldReturnStockAndOk(t *testing.T) {
	stockUsecase := new(StockUseCaseMock)

	stock := domain.Stock{
		Ticker: "UYA2",
		Price:  100,
	}

	stockUsecase.On("SearchByTicker", stock.Ticker).Return(stock, nil)

	controller := StockController{
		StockUseCase: stockUsecase,
	}

	foundStock, status := controller.FindByTicker(stock.Ticker)

	assert.Equal(t, status, http.StatusOK)
	assert.Equal(t, stock, foundStock)
}

func TestSearchUnregisteredStockShouldReturnNotFound(t *testing.T) {
	stockUsecase := new(StockUseCaseMock)

	stock := domain.Stock{
		Ticker: "UYA2",
		Price:  100,
	}

	error := domain.StockNotFoundError{Ticker: stock.Ticker}
	stockUsecase.On("SearchByTicker", stock.Ticker).Return(domain.Stock{}, error)

	controller := StockController{
		StockUseCase: stockUsecase,
	}

	_, status := controller.FindByTicker(stock.Ticker)

	assert.Equal(t, status, http.StatusNotFound)
}

func TestSearchStockWithUnexpectedErrorShouldReturnInternalServerError(t *testing.T) {
	stockUsecase := new(StockUseCaseMock)

	stock := domain.Stock{
		Ticker: "UYA2",
		Price:  100,
	}

	error := errors.New("Unexpected error")
	stockUsecase.On("SearchByTicker", stock.Ticker).Return(domain.Stock{}, error)

	controller := StockController{
		StockUseCase: stockUsecase,
	}

	_, status := controller.FindByTicker(stock.Ticker)

	assert.Equal(t, status, http.StatusInternalServerError)
}
