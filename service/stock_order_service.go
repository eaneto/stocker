package service

import "github.com/eaneto/stocker/repository"

type StockOrderService interface{}

type stockOrderService struct {
	StockOrderRepository repository.StockOrderRepository
}
