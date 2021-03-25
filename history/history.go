package history

import "time"

type StockHistoryEntity struct {
	ID        uint
	StockID   uint
	Price     uint
	CreatedAt time.Time
}
