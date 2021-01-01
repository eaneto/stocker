package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
)

var stockId uint = 0
var stockIdMutex sync.Mutex = sync.Mutex{}
var stocksMutex sync.Mutex = sync.Mutex{}
var stocksByTicker map[string]domain.StockEntity = make(map[string]domain.StockEntity)

type BaseStockRepository interface {
	Save(stock domain.StockEntity) error
	FindByTicker(ticker string) (domain.StockEntity, error)
	FindAll() []domain.StockEntity
}

type StockRepository struct{}

func NewStockRepository() BaseStockRepository {
	return StockRepository{}
}

func (repo StockRepository) Save(stock domain.StockEntity) error {
	stockIdMutex.Lock()
	stockId = stockId + 1
	stockIdMutex.Unlock()

	stocksMutex.Lock()
	stock.ID = stockId
	stocksByTicker[stock.Ticker] = stock
	stocksMutex.Unlock()
	return nil
}

func (repo StockRepository) FindByTicker(ticker string) (domain.StockEntity, error) {
	stock, ok := stocksByTicker[ticker]
	if ok {
		return stock, nil
	}
	return domain.StockEntity{}, domain.StockNotFoundError{Ticker: ticker}
}

func (repo StockRepository) FindAll() []domain.StockEntity {
	stocks := make([]domain.StockEntity, len(stocksByTicker))
	i := 0
	for _, stock := range stocksByTicker {
		stocks[i] = stock
		i = i + 1
	}
	return stocks
}
