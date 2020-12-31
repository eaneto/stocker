package domain

import (
	"math/big"
	"time"
)

const (
	STOCK        = "STOCK"
	STOCK_HISTOY = "STOCK"
)

type Stock struct {
	Ticker string
	Price  uint
}

type StockEntity struct {
	ID        big.Int
	Ticker    string
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StockHistoryEntity struct {
	ID        big.Int
	StockID   uint
	Price     uint
	CreatedAt time.Time
}
