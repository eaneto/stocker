package usecase

import (
	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/service"
)

type BaseStockUseCase interface {
	RegisterStock(stock domain.Stock) error
	SearchByTicker(ticker string) (domain.Stock, error)
}

type StockUseCase struct {
	StockService service.BaseStockService
}

func (usecase StockUseCase) RegisterStock(stock domain.Stock) error {
	usecase.StockService.RegisterStock(stock)
	return nil
}

func (usecase StockUseCase) SearchByTicker(ticker string) (domain.Stock, error) {
	return usecase.StockService.SearchByTicker(ticker)
}
