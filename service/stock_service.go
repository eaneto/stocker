package service

import (
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
)

type BaseStockService interface {
	RegisterStock(stock domain.Stock) error
	SearchByTicker(ticker string) (domain.Stock, error)
}

type StockService struct {
	StockRepository repository.BaseStockRepository
}

func (service StockService) RegisterStock(stock domain.Stock) error {
	stockEntity := domain.StockEntity{
		Ticker:    stock.Ticker,
		Price:     stock.Price,
		CreatedAt: time.Now(),
	}
	return service.StockRepository.Save(stockEntity)
}

func (service StockService) SearchByTicker(ticker string) (domain.Stock, error) {
	stockEntity, err := service.StockRepository.FindByTicker(ticker)
	return domain.Stock{
		Ticker: stockEntity.Ticker,
		Price:  stockEntity.Price,
	}, err
}
