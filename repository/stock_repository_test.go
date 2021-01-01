package repository

import (
	"testing"

	"github.com/eaneto/stocker/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveStock(t *testing.T) {
	clearAll()

	repository := StockRepository{}

	stock := domain.StockEntity{
		Ticker: "ABC",
	}

	err := repository.Save(stock)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(stocksByTicker))
	assert.Equal(t, stock.Ticker, stocksByTicker[stock.Ticker].Ticker)
	assert.Equal(t, uint(1), stocksByTicker[stock.Ticker].ID)
}

func TestFindStockByTicker(t *testing.T) {
	clearAll()

	stocks := make(map[string]domain.StockEntity)
	stock := domain.StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	stocks[stock.Ticker] = stock
	stocksByTicker = stocks

	repository := StockRepository{}

	foundStock, err := repository.FindByTicker(stock.Ticker)

	assert.Nil(t, err)
	assert.Equal(t, stock.Ticker, foundStock.Ticker)
	assert.Equal(t, stock.ID, foundStock.ID)
}

func TestFindStockByTickerNonExistentStock(t *testing.T) {
	clearAll()

	stock := domain.StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	repository := StockRepository{}

	_, err := repository.FindByTicker(stock.Ticker)

	assert.NotNil(t, err)
}
