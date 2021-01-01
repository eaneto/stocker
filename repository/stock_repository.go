package repository

import (
	"fmt"
	"sync"

	"github.com/eaneto/stocker/domain"
)

type StockNotFoundError struct {
	Ticker string
}

func (err StockNotFoundError) Error() string {
	return fmt.Sprintf("Stock not found, ticker=%s", err.Ticker)
}

var id uint = 0

type BaseStockRepository interface {
	Save(stock domain.StockEntity) error
	FindByTicker(ticker string) (domain.StockEntity, error)
	FindAll() []domain.StockEntity
}

type StockRepository struct {
	mu     sync.Mutex
	Stocks []domain.StockEntity
}

func (repo StockRepository) Save(stock domain.StockEntity) error {
	repo.mu.Lock()
	id = id + 1
	stock.ID = id
	repo.mu.Unlock()
	repo.Stocks = append(repo.Stocks, stock)
	return nil
}

func (repo StockRepository) FindByTicker(ticker string) (domain.StockEntity, error) {
	for _, stock := range repo.Stocks {
		if stock.Ticker == ticker {
			return stock, nil
		}
	}
	return domain.StockEntity{}, StockNotFoundError{Ticker: ticker}
}

func (repo StockRepository) FindAll() []domain.StockEntity {
	return repo.Stocks
}
