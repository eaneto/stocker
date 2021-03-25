package stock

import (
	"time"

	"github.com/sirupsen/logrus"
)

type BaseStockService interface {
	RegisterStock(stock Stock) error
	SearchByTicker(ticker string) (StockEntity, error)
	FindAll() []Stock
	FindByID(id uint) (StockEntity, error)
}

type StockService struct {
	StockRepository StockRepository
}

func NewStockService() BaseStockService {
	return StockService{
		StockRepository: NewStockRepository(),
	}
}

func (service StockService) RegisterStock(stock Stock) error {
	_, err := service.SearchByTicker(stock.Ticker)
	_, notFound := err.(StockNotFoundError)
	// If there are no stocks registered with the ticker.
	if notFound {
		return service.save(stock)
	}
	logrus.WithFields(logrus.Fields{
		"ticker": stock.Ticker,
	}).Warn("Already registered stock.")
	return AlreadyRegisteredStockError{Ticker: stock.Ticker}
}

func (service StockService) save(stock Stock) error {
	stockEntity := StockEntity{
		Ticker:    stock.Ticker,
		Price:     stock.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return service.StockRepository.Save(stockEntity)
}

func (service StockService) SearchByTicker(ticker string) (StockEntity, error) {
	return service.StockRepository.FindByTicker(ticker)
}

func (service StockService) FindByID(id uint) (StockEntity, error) {
	return service.StockRepository.FindByID(id)
}

func (service StockService) FindAll() []Stock {
	stocksEntities := service.StockRepository.FindAll()
	stocks := make([]Stock, len(stocksEntities))
	for i, entity := range stocksEntities {
		stocks[i] = Stock{
			Ticker: entity.Ticker,
			Price:  entity.Price,
		}
	}
	return stocks
}
