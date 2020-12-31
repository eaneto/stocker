package usecase

import (
	"github.com/eaneto/stocker/service"
	"github.com/google/uuid"
)

type StockOrderUseCase interface {
	CreateOrder(ticker string, customer uuid.UUID, amount int)
}

type orderUseCase struct {
	OrderService service.StockOrderService
}
