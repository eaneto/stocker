package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
)

var stockId uint = 0
var stockIdMutex sync.RWMutex = sync.RWMutex{}
var stocksMutex sync.RWMutex = sync.RWMutex{}

// stocksByTicker Stocks by ticker.
var stocksByTicker map[string]domain.StockEntity = make(map[string]domain.StockEntity)

// stocksByID Stocks by stock id.
var stocksByID map[uint]domain.StockEntity = make(map[uint]domain.StockEntity)

type StockRepository interface {
	Save(stock domain.StockEntity) error
	FindByTicker(ticker string) (domain.StockEntity, error)
	FindAll() []domain.StockEntity
	FindByID(id uint) (domain.StockEntity, error)
}

type StockRepositoryInMemory struct{}

func NewStockRepository() StockRepository {
	return StockRepositoryInMemory{}
}

func (StockRepositoryInMemory) Save(stock domain.StockEntity) error {
	stockIdMutex.Lock()
	stockId = stockId + 1
	stockIdMutex.Unlock()

	stocksMutex.Lock()
	stock.ID = stockId
	stocksByTicker[stock.Ticker] = stock
	stocksByID[stock.ID] = stock
	stocksMutex.Unlock()
	return nil
}

func (StockRepositoryInMemory) FindByTicker(ticker string) (domain.StockEntity, error) {
	stock, ok := stocksByTicker[ticker]
	if ok {
		return stock, nil
	}
	return domain.StockEntity{}, domain.StockNotFoundError{Ticker: ticker}
}

func (StockRepositoryInMemory) FindAll() []domain.StockEntity {
	stocks := make([]domain.StockEntity, len(stocksByTicker))
	i := 0
	for _, stock := range stocksByTicker {
		stocks[i] = stock
		i = i + 1
	}
	return stocks
}

func (StockRepositoryInMemory) FindByID(id uint) (domain.StockEntity, error) {
	stock, ok := stocksByID[id]
	if ok {
		return stock, nil
	}
	return domain.StockEntity{}, domain.StockNotFoundError{ID: id}
}
