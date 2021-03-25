package stock

import (
	"fmt"
	"time"
)

const (
	STOCK        = "STOCK"
	STOCK_HISTOY = "STOCK"
)

type AlreadyRegisteredStockError struct {
	Ticker string
}

func (e AlreadyRegisteredStockError) Error() string {
	return fmt.Sprintf("Stock already registered with ticker=%s", e.Ticker)
}

type StockNotFoundError struct {
	Ticker string
	ID     uint
}

func (err StockNotFoundError) Error() string {
	return fmt.Sprintf("Stock not found, ticker=%s, id=%d", err.Ticker, err.ID)
}

type Stock struct {
	Ticker string `json:"ticker"`
	Price  uint   `json:"price"`
}

type StockEntity struct {
	ID        uint
	Ticker    string
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
