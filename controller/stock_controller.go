package controller

import (
	"net/http"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/usecase"
)

type BaseStockController interface {
	RegisterStock(stock domain.Stock) (httpStatus int)
	FindByTicker(ticker string) (stock domain.Stock, httpStatus int)
	FindAll() (stocks []domain.Stock, httpStatus int)
}

type StockController struct {
	StockUseCase usecase.BaseStockUseCase
}

func NewStockController() BaseStockController {
	return StockController{
		StockUseCase: usecase.NewStockUseCase(),
	}
}

func (controller StockController) RegisterStock(stock domain.Stock) (httpStatus int) {
	err := controller.StockUseCase.RegisterStock(stock)
	if err == nil {
		return http.StatusCreated
	}
	_, isConflicError := err.(domain.AlreadyRegisteredStockError)
	if isConflicError {
		return http.StatusConflict
	} else {
		return http.StatusInternalServerError
	}
}

func (controller StockController) FindByTicker(ticker string) (domain.Stock, int) {
	stock, err := controller.StockUseCase.SearchByTicker(ticker)
	if err == nil {
		return stock, http.StatusOK
	}
	_, isNotFound := err.(domain.StockNotFoundError)
	if isNotFound {
		return domain.Stock{}, http.StatusNotFound
	}
	return domain.Stock{}, http.StatusInternalServerError
}

func (controller StockController) FindAll() ([]domain.Stock, int) {
	return controller.StockUseCase.FindAll(), http.StatusOK
}
