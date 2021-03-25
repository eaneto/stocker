package order

import (
	"sync"

	"github.com/google/uuid"
)

var stockOrderMutex sync.RWMutex = sync.RWMutex{}

// Orders by the order code.
var ordersByCode map[uuid.UUID]StockOrderEntity = make(map[uuid.UUID]StockOrderEntity)

// Orders by the customer id.
var ordersByCustomer map[uint][]StockOrderEntity = make(map[uint][]StockOrderEntity)

type StockOrderRepository interface {
	Save(stockOrder StockOrderEntity) error
	Update(stockOrder StockOrderEntity) error
	FindByCode(code uuid.UUID) (StockOrderEntity, error)
	FindAllByCustomer(customerID uint) []StockOrderEntity
}

type StockOrderRepositoryInMemory struct{}

func NewStockOrderRepository() StockOrderRepository {
	return StockOrderRepositoryInMemory{}
}

func (StockOrderRepositoryInMemory) Save(stockOrder StockOrderEntity) error {
	stockOrderMutex.Lock()
	ordersByCode[stockOrder.Code] = stockOrder
	ordersByCustomer[stockOrder.CustomerID] = append(
		ordersByCustomer[stockOrder.CustomerID], stockOrder)
	stockOrderMutex.Unlock()
	return nil
}

func (StockOrderRepositoryInMemory) Update(stockOrder StockOrderEntity) error {
	stockOrderMutex.Lock()
	ordersByCode[stockOrder.Code] = stockOrder
	// Updates the order on the customers list.
	for index, order := range ordersByCustomer[stockOrder.CustomerID] {
		if order.Code == order.Code {
			ordersByCustomer[stockOrder.CustomerID][index] = order
		}
	}
	stockOrderMutex.Unlock()
	return nil
}

func (StockOrderRepositoryInMemory) FindByCode(code uuid.UUID) (StockOrderEntity, error) {
	order, ok := ordersByCode[code]
	if ok {
		return order, nil
	}
	return StockOrderEntity{}, StockOrderNotFoundError{Code: code}
}

func (StockOrderRepositoryInMemory) FindAllByCustomer(customerID uint) []StockOrderEntity {
	return ordersByCustomer[customerID]
}
