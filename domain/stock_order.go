package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const STOCK_ORDER = "stock_order"

type StockOrderNotFoundError struct {
	Code uuid.UUID
}

func (e StockOrderNotFoundError) Error() string {
	return fmt.Sprintf("Stock order not found. code=%s", e.Code)
}

type OrderStatus string

const (
	Created   OrderStatus = "CREATED"
	Confirmed OrderStatus = "CONFIRMED"
	Cancelled OrderStatus = "CANCELLED"
)

type StockOrderEntity struct {
	Code       uuid.UUID
	CustomerID uint
	StockID    uint
	Amount     uint
	Status     OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type StockOrderRequest struct {
	Code        uuid.UUID `json:"code"`
	CustomerID  uint      `json:"customer_id"`
	StockTicker string    `json:"stock_ticker"`
	Amount      uint      `json:"amount"`
}

type StockPosition struct {
	Ticker string `json:"ticker"`
	Price  uint   `json:"price"`
	Amount uint   `json:"amount"`
}

type CustomerPosition struct {
	CustomerID uint            `json:"customer_id"`
	Stocks     []StockPosition `json:"stock"`
}
