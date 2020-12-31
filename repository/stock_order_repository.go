package repository

import "github.com/eaneto/stocker/domain"

type StockOrderRepository interface {
	Save(stockOrder domain.StockOrderEntity) error
}
