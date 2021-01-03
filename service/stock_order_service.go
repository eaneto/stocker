package service

import (
	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
	"github.com/google/uuid"
)

type BaseStockOrderService interface {
	CreateOrder(stockOrderRequest domain.StockOrderRequest) error
	FindAllByCustomer(customerID uint) []domain.StockOrderEntity
	GetCustomerPosition(customerCode uuid.UUID) (domain.CustomerPosition, error)
}

type StockOrderService struct {
	StockOrderRepository repository.StockOrderRepository
	CustomerService      BaseCustomerService
	StockService         BaseStockService
}

func (service StockOrderService) GetCustomerPosition(customerCode uuid.UUID) (domain.CustomerPosition, error) {
	// Find customer id by customerCode
	customer, err := service.CustomerService.FindByCode(customerCode)
	_, notFound := err.(domain.CustomerNotFoundError)
	if notFound {
		return domain.CustomerPosition{}, err
	}
	// Get all orders from customer id
	orders := service.StockOrderRepository.FindAllByCustomer(customer.ID)
	// Map orders to domain.CustomerPosition
	orderAmountByID := make(map[uint]uint)
	for _, order := range orders {
		orderAmountByID[order.StockID] = orderAmountByID[order.StockID] + order.Amount
	}
	// Ignore if stock not found because it's an impossible case
	stocks := make([]domain.StockPosition, len(orderAmountByID))
	for id, amount := range orderAmountByID {
		stock, _ := service.StockService.FindByID(id)
		stockPosition := domain.StockPosition{
			Ticker: stock.Ticker,
			Price:  stock.Price,
			Amount: amount,
		}
		stocks = append(stocks, stockPosition)
	}
	position := domain.CustomerPosition{
		CustomerCode: customerCode,
		Stocks:       stocks,
	}
	return position, nil
}
