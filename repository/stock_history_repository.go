package repository

import "github.com/eaneto/stocker/domain"

type BaseStockHistoryRepository interface {
	Save(stock domain.StockHistoryEntity) error
}
