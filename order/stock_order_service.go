package order

import (
	"time"

	"github.com/eaneto/stocker/stock"
	"github.com/sirupsen/logrus"
)

var orderRequestsChannel chan StockOrderRequest = make(chan StockOrderRequest)

type BaseStockOrderService interface {
	// CreateOrder Creates an order.
	CreateOrder(stockOrderRequest StockOrderRequest) error
	// ConfirmOrder Confirms an order.
	ConfirmOrders()
	FindAllByCustomer(customerID uint) []StockOrderEntity
	// GetCustomerPosition Get customer position on all stocks.
	GetCustomerPosition(customerID uint) (CustomerPosition, error)
}

type StockOrderService struct {
	StockOrderRepository StockOrderRepository
	StockService         stock.BaseStockService
}

func NewStockOrderService() BaseStockOrderService {
	return StockOrderService{
		StockOrderRepository: NewStockOrderRepository(),
		StockService:         stock.NewStockService(),
	}
}

func (service StockOrderService) CreateOrder(stockOrderRequest StockOrderRequest) error {
	stock, err := service.StockService.SearchByTicker(stockOrderRequest.StockTicker)
	if err != nil {
		return err
	}

	_, err = service.StockOrderRepository.FindByCode(stockOrderRequest.Code)
	_, notFound := err.(StockOrderNotFoundError)
	if !notFound {
		return StockOrderAlreadyProcessedError{Code: stockOrderRequest.Code}
	}

	stockOrder := StockOrderEntity{
		CustomerID: stockOrderRequest.CustomerID,
		StockID:    stock.ID,
		Amount:     stockOrderRequest.Amount,
		Code:       stockOrderRequest.Code,
		Status:     Created,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	service.StockOrderRepository.Save(stockOrder)
	// Publish the request to the stock order channel
	orderRequestsChannel <- stockOrderRequest
	return nil
}

func (service StockOrderService) ConfirmOrders() {
	for {
		select {
		case orderRequest := <-orderRequestsChannel:
			// Expensive computation
			time.Sleep(time.Second * 2)
			service.confirmOrder(orderRequest)
		}
	}
}

func (service StockOrderService) confirmOrder(stockOrderRequest StockOrderRequest) {
	order, err := service.StockOrderRepository.FindByCode(stockOrderRequest.Code)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"order_code": stockOrderRequest.Code,
			"error":      err,
		}).Error("Reading order, this should never happen!!")
	} else {
		order.Status = Confirmed
		order.UpdatedAt = time.Now()
		service.StockOrderRepository.Update(order)
		logrus.WithField("order_code", order.Code).
			Info("Order confirmed succcessfully")
	}
}

func (service StockOrderService) FindAllByCustomer(customerID uint) []StockOrderEntity {
	return nil
}

func (service StockOrderService) GetCustomerPosition(customerID uint) (CustomerPosition, error) {
	// Get all orders from customer id
	orders := service.StockOrderRepository.FindAllByCustomer(customerID)
	orderAmountByID := summarizeOrdersByStock(orders)
	stocks := service.mapOrdersToStockPosition(orderAmountByID)
	position := CustomerPosition{
		CustomerID: customerID,
		Stocks:     stocks,
	}
	return position, nil
}

// mapOrdersToStockPosition Maps a map of the amounts by the stock id
// to a list of StockPosition.
func (service StockOrderService) mapOrdersToStockPosition(orderAmountByID map[uint]uint) []StockPosition {
	i := 0
	stocks := make([]StockPosition, len(orderAmountByID))
	for id, amount := range orderAmountByID {
		// Ignore if stock not found because it's an impossible case
		stock, _ := service.StockService.FindByID(id)
		stockPosition := StockPosition{
			Ticker: stock.Ticker,
			Price:  stock.Price,
			Amount: amount,
		}
		stocks[i] = stockPosition
		i++
	}
	return stocks
}

// summarizeOrdersByStock Summarize a list of orders by the stock id.
func summarizeOrdersByStock(orders []StockOrderEntity) map[uint]uint {
	orderAmountByID := make(map[uint]uint)
	for _, order := range orders {
		orderAmountByID[order.StockID] = orderAmountByID[order.StockID] + order.Amount
	}
	return orderAmountByID
}
