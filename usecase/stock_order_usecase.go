package usecase

import (
	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/service"
	"github.com/google/uuid"
)

type BaseStockOrderUseCase interface {
	CreateOrder(ticker string, customerCode uuid.UUID, amount int)
	GetCustomerPosition(customerCode uuid.UUID) (domain.CustomerPosition, error)
}

type StockOrderUseCase struct {
	StockService      service.BaseStockService
	CustomerService   service.BaseCustomerService
	StockOrderService service.BaseStockOrderService
}

func (usecase StockOrderUseCase) GetCustomerPosition(customerCode uuid.UUID) (domain.CustomerPosition, error) {
	// Find customer id by customerCode
	customer, err := usecase.CustomerService.FindByCode(customerCode)
	_, notFound := err.(domain.CustomerNotFoundError)
	if notFound {
		return domain.CustomerPosition{}, err
	}
	// Get all orders from customer id
	orders := usecase.StockOrderService.FindAllByCustomer(customer.ID)
	// Map orders to domain.CustomerPosition
	orderAmountByID := make(map[uint]uint)
	for _, order := range orders {
		orderAmountByID[order.StockID] = orderAmountByID[order.StockID] + order.Amount
	}
	// Ignore if stock not found because it's an impossible case
	stocks := make([]domain.StockPosition, len(orderAmountByID))
	for id, amount := range orderAmountByID {
		stock, _ := usecase.StockService.FindByID(id)
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
