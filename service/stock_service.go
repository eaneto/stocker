package service

import (
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
	"github.com/sirupsen/logrus"
)

type BaseStockService interface {
	RegisterStock(stock domain.Stock) error
	SearchByTicker(ticker string) (domain.Stock, error)
	FindAll() []domain.Stock
	FindByID(id uint) (domain.StockEntity, error)
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
	_, err := service.SearchByTicker(stock.Ticker)
	_, notFound := err.(domain.StockNotFoundError)
	// If there are no stocks registered with the ticker.
	if notFound {
		return service.save(stock)
	}
	logrus.WithFields(logrus.Fields{
		"ticker": stock.Ticker,
	}).Warn("Already registered stock.")
	return domain.AlreadyRegisteredStockError{Ticker: stock.Ticker}
}

func (service StockService) save(stock domain.Stock) error {
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

func (service StockService) FindByID(id uint) (domain.StockEntity, error) {
	return service.StockRepository.FindByID(id)
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
