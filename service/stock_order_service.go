package service

import (
	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
)

type BaseStockOrderService interface {
	CreateOrder(stockOrderRequest domain.StockOrderRequest) error
	FindAllByCustomer(customerID uint) []domain.StockOrderEntity
}

type StockOrderService struct {
	StockOrderRepository repository.StockOrderRepository
}
