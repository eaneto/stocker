package service

import (
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
)

type BaseStockService interface {
	RegisterStock(stock domain.Stock) error
	SearchByTicker(ticker string) (domain.Stock, error)
	FindAll() []domain.Stock
}

type StockService struct {
	StockRepository repository.BaseStockRepository
}

func NewStockService() BaseStockService {
	return StockService{
		StockRepository: repository.NewStockRepository(),
	}
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

func (service StockService) FindAll() []domain.Stock {
	stocksEntities := service.StockRepository.FindAll()
	stocks := make([]domain.Stock, len(stocksEntities))
	for i, entity := range stocksEntities {
		stocks[i] = domain.Stock{
			Ticker: entity.Ticker,
			Price:  entity.Price,
		}
	}
	return stocks
}
