package stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveStockShouldSaveStockOnMap(t *testing.T) {
	clearAllStocks()

	repository := StockRepositoryInMemory{}

	stock := StockEntity{
		Ticker: "ABC",
	}

	err := repository.Save(stock)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(stocksByTicker))
	assert.Equal(t, 1, len(stocksByID))
	assert.Equal(t, stock.Ticker, stocksByTicker[stock.Ticker].Ticker)
	assert.Equal(t, uint(1), stocksByTicker[stock.Ticker].ID)
}

func TestFindStockByTicker(t *testing.T) {
	clearAllStocks()

	stocks := make(map[string]StockEntity)
	stock := StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	stocks[stock.Ticker] = stock
	stocksByTicker = stocks

	repository := StockRepositoryInMemory{}

	foundStock, err := repository.FindByTicker(stock.Ticker)

	assert.Nil(t, err)
	assert.Equal(t, stock.Ticker, foundStock.Ticker)
	assert.Equal(t, stock.ID, foundStock.ID)
}

func TestFindStockByTickerNonExistentStock(t *testing.T) {
	clearAllStocks()

	stock := StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	repository := StockRepositoryInMemory{}

	_, err := repository.FindByTicker(stock.Ticker)

	assert.NotNil(t, err)
	_, notFound := err.(StockNotFoundError)
	assert.True(t, notFound)
}

func TestFindStockByID(t *testing.T) {
	clearAllStocks()

	stocks := make(map[string]StockEntity)
	stock := StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	stocks[stock.Ticker] = stock
	stocksByTicker = stocks

	repository := StockRepositoryInMemory{}

	foundStock, err := repository.FindByID(stock.ID)

	assert.Nil(t, err)
	assert.Equal(t, stock.Ticker, foundStock.Ticker)
	assert.Equal(t, stock.ID, foundStock.ID)
}

func TestFindStockByIDNonExistentStockShouldReturnError(t *testing.T) {
	clearAllStocks()

	stock := StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	repository := StockRepositoryInMemory{}

	_, err := repository.FindByTicker(stock.Ticker)

	assert.NotNil(t, err)
	_, notFound := err.(StockNotFoundError)
	assert.True(t, notFound)
}

func TestFindAllStocksWithNoneRegisteredShouldReturnEmptySlice(t *testing.T) {
	clearAllStocks()

	repository := StockRepositoryInMemory{}

	stocks := repository.FindAll()

	assert.Empty(t, stocks)
}

func TestFindAllStocksWithOneRegisteredShouldReturnSliceWithOneElement(t *testing.T) {
	clearAllStocks()

	stocks := make(map[string]StockEntity)
	stock := StockEntity{
		ID:     1,
		Ticker: "ABC",
	}
	stocks[stock.Ticker] = stock
	stocksByTicker = stocks

	repository := StockRepositoryInMemory{}

	foundStocks := repository.FindAll()

	assert.Equal(t, len(stocks), len(foundStocks))
	assert.Equal(t, stock, foundStocks[0])
}

// clearAllStocks Clears all stored stocks and resets id.
func clearAllStocks() {
	stockId = 0
	stocksByTicker = make(map[string]StockEntity)
}
