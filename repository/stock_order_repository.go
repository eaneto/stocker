package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
)

var stockOrderId uint = 0
var stockOrderIdMutex sync.Mutex = sync.Mutex{}
var stockOrderMutex sync.Mutex = sync.Mutex{}

// Orders by the order code.
var ordersByCode map[uuid.UUID]domain.StockOrderEntity = make(map[uuid.UUID]domain.StockOrderEntity)

// Orders by the customer id.
var ordersByCustomer map[uint][]domain.StockOrderEntity = make(map[uint][]domain.StockOrderEntity)

type BaseStockOrderRepository interface {
	Save(stockOrder domain.StockOrderEntity) error
}

type StockOrderRepository struct{}

func (StockOrderRepository) Save(stockOrder domain.StockOrderEntity) error {
	stockOrderIdMutex.Lock()
	stockOrderId = stockOrderId + 1
	stockOrderIdMutex.Unlock()

	stockOrder.ID = stockOrderId

	stockOrderMutex.Lock()
	ordersByCode[stockOrder.Code] = stockOrder
	ordersByCustomer[stockOrder.CustomerID] = append(ordersByCustomer[stockOrder.CustomerID], stockOrder)
	stockOrderMutex.Unlock()
	return nil
}
