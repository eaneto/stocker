package domain

import (
	"time"
)

const (
	STOCK        = "STOCK"
	STOCK_HISTOY = "STOCK"
)

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
