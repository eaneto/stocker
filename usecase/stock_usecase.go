package usecase

import (
	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/service"
	"github.com/sirupsen/logrus"
)

type BaseStockUseCase interface {
	RegisterStock(stock domain.Stock) error
	SearchByTicker(ticker string) (domain.Stock, error)
}

type StockUseCase struct {
	StockService service.BaseStockService
}

func NewStockUseCase() BaseStockUseCase {
	return StockUseCase{
		StockService: service.NewStockService(),
	}
}

func (usecase StockUseCase) RegisterStock(stock domain.Stock) error {
	_, err := usecase.SearchByTicker(stock.Ticker)
	_, notFound := err.(domain.StockNotFoundError)
	// If there are no stocks registered with the ticker.
	if notFound {
		return usecase.StockService.RegisterStock(stock)
	}
	logrus.WithFields(logrus.Fields{
		"ticker": stock.Ticker,
	}).Warn("Already registered stock.")
	return domain.AlreadyRegisteredStockError{Ticker: stock.Ticker}
}

func (usecase StockUseCase) SearchByTicker(ticker string) (domain.Stock, error) {
	return usecase.StockService.SearchByTicker(ticker)
}
