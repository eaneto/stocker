package repository

import "github.com/eaneto/stocker/domain"

type StockRepository interface {
	Save(stock domain.StockEntity) error
	FindByTicker(ticker string) domain.StockEntity
	FindAll() []domain.StockEntity
}

type StockHistoryRepository interface {
	Save(stock domain.StockHistoryEntity) error
}
