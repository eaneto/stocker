package repository

import (
	"sync"
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveStock(t *testing.T) {
	repository := StockRepository{
		mu:     sync.Mutex{},
		Stocks: make(map[string]domain.StockEntity),
	}

	stock := domain.StockEntity{
		Ticker: "ABC",
	}

	err := repository.Save(stock)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(repository.Stocks))
	assert.Equal(t, stock.Ticker, repository.Stocks[stock.Ticker].Ticker)
	assert.Equal(t, uint(1), repository.Stocks[stock.Ticker].ID)
}

func TestFindStockByTicker(t *testing.T) {
	stocks := make(map[string]domain.StockEntity)
	stock := domain.StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	stocks[stock.Ticker] = stock
	repository := StockRepository{
		mu:     sync.Mutex{},
		Stocks: stocks,
	}

	foundStock, err := repository.FindByTicker(stock.Ticker)

	assert.Nil(t, err)
	assert.Equal(t, stock.Ticker, foundStock.Ticker)
	assert.Equal(t, stock.ID, foundStock.ID)
}

func TestFindStockByTickerNonExistentStock(t *testing.T) {
	stocks := make(map[string]domain.StockEntity)
	stock := domain.StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	repository := StockRepository{
		mu:     sync.Mutex{},
		Stocks: stocks,
	}

	_, err := repository.FindByTicker(stock.Ticker)

	assert.NotNil(t, err)
}
