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
