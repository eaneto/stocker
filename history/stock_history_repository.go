package history

type BaseStockHistoryRepository interface {
	Save(stock StockHistoryEntity) error
}
