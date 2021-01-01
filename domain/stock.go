package domain

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
}

func (err StockNotFoundError) Error() string {
	return fmt.Sprintf("Stock not found, ticker=%s", err.Ticker)
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

type StockHistoryEntity struct {
	ID        uint
	StockID   uint
	Price     uint
	CreatedAt time.Time
}
