package domain

import (
	"time"

	"github.com/google/uuid"
)

const STOCK_ORDER = "stock_order"

type OrderStatus string

const (
	Created   OrderStatus = "CREATED"
	Confirmed OrderStatus = "CONFIRMED"
	Cancelled OrderStatus = "CANCELLED"
)

type StockOrderEntity struct {
	ID         uint
	Code       uuid.UUID
	CustomerID uint
	StockID    uint
	Amount     uint
	Status     OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type StockOrderRequest struct {
	Code         uuid.UUID `json:"code"`
	CustomerCode uuid.UUID `json:"customer_code"`
	StockTicker  uint      `json:"stock_ticker"`
	Amount       uint      `json:"amount"`
}

type StockPosition struct {
	Ticker string `json:"ticker"`
	Price  uint   `json:"price"`
	Amount uint   `json:"amount"`
}

type CustomerPosition struct {
	CustomerCode uuid.UUID       `json:"customer_code"`
	Stocks       []StockPosition `json:"stock"`
}
