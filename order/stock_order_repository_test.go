package order

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveStockOrderShouldSaveByCodeAndCustomer(t *testing.T) {
	clearAllStockOrders()

	repository := StockOrderRepositoryInMemory{}
	code, _ := uuid.NewRandom()
	stockOrder := StockOrderEntity{
		Code:       code,
		CustomerID: uint(1),
		StockID:    uint(2),
		Amount:     uint(100),
		Status:     Created,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := repository.Save(stockOrder)

	assert.Nil(t, err)
	assert.NotEmpty(t, ordersByCode)
	assert.NotEmpty(t, ordersByCustomer)
	assert.NotEmpty(t, ordersByCustomer[stockOrder.CustomerID])
	foundOrderByCode := ordersByCode[stockOrder.Code]
	assertStockOrderEqual(t, stockOrder, foundOrderByCode)
	foundOrderByCustomer := ordersByCustomer[stockOrder.CustomerID][0]
	assertStockOrderEqual(t, stockOrder, foundOrderByCustomer)
}

func TestSaveTwoStockOrderFromSameCustomerShouldSaveByCodeAndCustomer(t *testing.T) {
	clearAllStockOrders()

	repository := StockOrderRepositoryInMemory{}
	code1, _ := uuid.NewRandom()
	code2, _ := uuid.NewRandom()
	customerID := uint(1)
	stockOrders := []StockOrderEntity{
		{
			Code:       code1,
			CustomerID: customerID,
			StockID:    uint(2),
			Amount:     uint(100),
			Status:     Created,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Code:       code2,
			CustomerID: customerID,
			StockID:    uint(4),
			Amount:     uint(50),
			Status:     Created,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	err := repository.Save(stockOrders[0])
	assert.Nil(t, err)
	err = repository.Save(stockOrders[1])
	assert.Nil(t, err)

	assert.NotEmpty(t, ordersByCode)
	assert.NotEmpty(t, ordersByCustomer)
	assert.NotEmpty(t, ordersByCustomer[customerID])

	expectedSize := 2
	assert.Equal(t, expectedSize, len(ordersByCode))
	assert.Equal(t, expectedSize, len(ordersByCustomer[customerID]))

	foundOrderByCode := ordersByCode[stockOrders[0].Code]
	assertStockOrderEqual(t, stockOrders[0], foundOrderByCode)
	foundOrderByCustomer := ordersByCustomer[customerID]
	assertStockOrderEqual(t, stockOrders[0], foundOrderByCustomer[0])
	assertStockOrderEqual(t, stockOrders[1], foundOrderByCustomer[1])
}

func assertStockOrderEqual(t *testing.T, orderA, orderB StockOrderEntity) {
	assert.Equal(t, orderA.Amount, orderB.Amount)
	assert.Equal(t, orderA.StockID, orderB.StockID)
	assert.Equal(t, orderA.Status, orderB.Status)
	assert.Equal(t, orderA.CreatedAt, orderB.CreatedAt)
	assert.Equal(t, orderA.UpdatedAt, orderB.UpdatedAt)
}

func TestFindAllStockOrdersByCustomer(t *testing.T) {
	clearAllStockOrders()

	repository := StockOrderRepositoryInMemory{}
	customerID := uint(1)

	orders := repository.FindAllByCustomer(customerID)

	assert.Empty(t, orders)
}

// clearAllStockOrders Clears all stored orders and resets id.
func clearAllStockOrders() {
	ordersByCode = make(map[uuid.UUID]StockOrderEntity)
	ordersByCustomer = make(map[uint][]StockOrderEntity)
}
