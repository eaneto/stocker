package stock

import (
	"sync"
)

var stockId uint = 0
var stockIdMutex sync.RWMutex = sync.RWMutex{}
var stocksMutex sync.RWMutex = sync.RWMutex{}

// stocksByTicker Stocks by ticker.
var stocksByTicker map[string]StockEntity = make(map[string]StockEntity)

// stocksByID Stocks by stock id.
var stocksByID map[uint]StockEntity = make(map[uint]StockEntity)

type StockRepository interface {
	Save(stock StockEntity) error
	FindByTicker(ticker string) (StockEntity, error)
	FindAll() []StockEntity
	FindByID(id uint) (StockEntity, error)
}

type StockRepositoryInMemory struct{}

func NewStockRepository() StockRepository {
	return StockRepositoryInMemory{}
}

func (StockRepositoryInMemory) Save(stock StockEntity) error {
	stockIdMutex.Lock()
	stockId++
	stockIdMutex.Unlock()

	stock.ID = stockId
	stocksMutex.Lock()
	stocksByTicker[stock.Ticker] = stock
	stocksByID[stock.ID] = stock
	stocksMutex.Unlock()
	return nil
}

func (StockRepositoryInMemory) FindByTicker(ticker string) (StockEntity, error) {
	stock, ok := stocksByTicker[ticker]
	if ok {
		return stock, nil
	}
	return StockEntity{}, StockNotFoundError{Ticker: ticker}
}

func (StockRepositoryInMemory) FindAll() []StockEntity {
	stocks := make([]StockEntity, len(stocksByTicker))
	i := 0
	for _, stock := range stocksByTicker {
		stocks[i] = stock
		i++
	}
	return stocks
}

func (StockRepositoryInMemory) FindByID(id uint) (StockEntity, error) {
	stock, ok := stocksByID[id]
	if ok {
		return stock, nil
	}
	return StockEntity{}, StockNotFoundError{ID: id}
}
