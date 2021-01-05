package service

import (
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/repository"
	"github.com/sirupsen/logrus"
)

var orderRequests chan domain.StockOrderRequest = make(chan domain.StockOrderRequest)

type BaseStockOrderService interface {
	// CreateOrder Creates an order.
	CreateOrder(stockOrderRequest domain.StockOrderRequest) error
	// ConfirmOrder Confirms an order.
	ConfirmOrders()
	FindAllByCustomer(customerID uint) []domain.StockOrderEntity
	// GetCustomerPosition Get customer position on all stocks.
	GetCustomerPosition(customerID uint) (domain.CustomerPosition, error)
}

type StockOrderService struct {
	StockOrderRepository repository.StockOrderRepository
	StockService         BaseStockService
}

func NewStockOrderService() BaseStockOrderService {
	return StockOrderService{
		StockOrderRepository: repository.NewStockOrderRepository(),
		StockService:         NewStockService(),
	}
}

func (service StockOrderService) CreateOrder(stockOrderRequest domain.StockOrderRequest) error {
	stock, err := service.StockService.SearchByTicker(stockOrderRequest.StockTicker)
	if err != nil {
		return err
	}

	_, err = service.StockOrderRepository.FindByCode(stockOrderRequest.Code)
	_, notFound := err.(domain.StockOrderNotFoundError)
	if !notFound {
		return domain.StockOrderAlreadyProcessedError{Code: stockOrderRequest.Code}
	}

	stockOrder := domain.StockOrderEntity{
		CustomerID: stockOrderRequest.CustomerID,
		StockID:    stock.ID,
		Amount:     stockOrderRequest.Amount,
		Code:       stockOrderRequest.Code,
		Status:     domain.Created,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	service.StockOrderRepository.Save(stockOrder)
	// Publish the request to the stock order channel
	orderRequests <- stockOrderRequest
	return nil
}

func (service StockOrderService) ConfirmOrders() {
	for {
		select {
		case orderRequest := <-orderRequests:
			service.confirmOrder(orderRequest)
		}
	}
}

func (service StockOrderService) confirmOrder(stockOrderRequest domain.StockOrderRequest) {
	order, err := service.StockOrderRepository.FindByCode(stockOrderRequest.Code)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"order_code": stockOrderRequest.Code,
			"error":      err,
		}).Error("Reading order, this should never happen!!")
	} else {
		order.Status = domain.Confirmed
		order.UpdatedAt = time.Now()
		service.StockOrderRepository.Update(order)
		logrus.WithField("order_code", order.Code).
			Info("Order confirmed succcessfully")
	}
}

func (service StockOrderService) FindAllByCustomer(customerID uint) []domain.StockOrderEntity {
	return nil
}

func (service StockOrderService) GetCustomerPosition(customerID uint) (domain.CustomerPosition, error) {
	// Get all orders from customer id
	orders := service.StockOrderRepository.FindAllByCustomer(customerID)
	// Map orders to domain.CustomerPosition
	orderAmountByID := make(map[uint]uint)
	for _, order := range orders {
		orderAmountByID[order.StockID] = orderAmountByID[order.StockID] + order.Amount
	}
	stocks := make([]domain.StockPosition, len(orderAmountByID))
	// Loop through all orders to summarize the orders into a stock
	// position.
	i := 0
	for id, amount := range orderAmountByID {
		// Ignore if stock not found because it's an impossible case
		stock, _ := service.StockService.FindByID(id)
		stockPosition := domain.StockPosition{
			Ticker: stock.Ticker,
			Price:  stock.Price,
			Amount: amount,
		}
		stocks[i] = stockPosition
		i++
	}
	position := domain.CustomerPosition{
		CustomerID: customerID,
		Stocks:     stocks,
	}
	return position, nil
}
