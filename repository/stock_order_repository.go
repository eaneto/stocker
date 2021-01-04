package repository

import (
	"sync"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
)

var stockOrderMutex sync.RWMutex = sync.RWMutex{}

// Orders by the order code.
var ordersByCode map[uuid.UUID]domain.StockOrderEntity = make(map[uuid.UUID]domain.StockOrderEntity)

// Orders by the customer id.
var ordersByCustomer map[uint][]domain.StockOrderEntity = make(map[uint][]domain.StockOrderEntity)

type StockOrderRepository interface {
	Save(stockOrder domain.StockOrderEntity) error
	Update(stockOrder domain.StockOrderEntity) error
	FindByCode(code uuid.UUID) (domain.StockOrderEntity, error)
	FindAllByCustomer(customerID uint) []domain.StockOrderEntity
}

type StockOrderRepositoryInMemory struct{}

func NewStockOrderRepository() StockOrderRepository {
	return StockOrderRepositoryInMemory{}
}

func (StockOrderRepositoryInMemory) Save(stockOrder domain.StockOrderEntity) error {
	stockOrderMutex.Lock()
	ordersByCode[stockOrder.Code] = stockOrder
	ordersByCustomer[stockOrder.CustomerID] = append(
		ordersByCustomer[stockOrder.CustomerID], stockOrder)
	stockOrderMutex.Unlock()
	return nil
}

func (StockOrderRepositoryInMemory) Update(stockOrder domain.StockOrderEntity) error {
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

func (StockOrderRepositoryInMemory) FindByCode(code uuid.UUID) (domain.StockOrderEntity, error) {
	order, ok := ordersByCode[code]
	if ok {
		return order, nil
	}
	return domain.StockOrderEntity{}, domain.StockOrderNotFoundError{Code: code}
}

func (StockOrderRepositoryInMemory) FindAllByCustomer(customerID uint) []domain.StockOrderEntity {
	return ordersByCustomer[customerID]
}
