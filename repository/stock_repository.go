package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
)

var id uint = 0
var mu sync.Mutex = sync.Mutex{}
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
	mu.Lock()
	id = id + 1
	stock.ID = id
	mu.Unlock()
	stocksByTicker[stock.Ticker] = stock
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
	for _, stock := range stocksByTicker {
		stocks = append(stocks, stock)
	}
	return stocks
}

// clearAll Clears all stored data, meant to used only on tests.
func clearAll() {
	id = 0
	stocksByTicker = make(map[string]domain.StockEntity)
}
