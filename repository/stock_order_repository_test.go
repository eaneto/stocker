package repository

import (
	"testing"
	"time"

	"github.com/eaneto/stocker/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveStockOrderShouldSaveByCodeAndCustomer(t *testing.T) {
	clearAllStockOrders()

	repository := StockOrderRepository{}
	code, _ := uuid.NewRandom()
	stockOrder := domain.StockOrderEntity{
		Code:       code,
		CustomerID: uint(1),
		StockID:    uint(2),
		Amount:     uint(100),
		Status:     domain.Created,
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

	repository := StockOrderRepository{}
	code1, _ := uuid.NewRandom()
	code2, _ := uuid.NewRandom()
	customerID := uint(1)
	stockOrders := []domain.StockOrderEntity{
		{
			Code:       code1,
			CustomerID: customerID,
			StockID:    uint(2),
			Amount:     uint(100),
			Status:     domain.Created,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			Code:       code2,
			CustomerID: customerID,
			StockID:    uint(4),
			Amount:     uint(50),
			Status:     domain.Created,
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

func assertStockOrderEqual(t *testing.T, orderA, orderB domain.StockOrderEntity) {
	assert.NotEqual(t, uint(0), orderB.ID)
	assert.Equal(t, orderA.Amount, orderB.Amount)
	assert.Equal(t, orderA.StockID, orderB.StockID)
	assert.Equal(t, orderA.Status, orderB.Status)
	assert.Equal(t, orderA.CreatedAt, orderB.CreatedAt)
	assert.Equal(t, orderA.UpdatedAt, orderB.UpdatedAt)
}

// clearAllStockOrders Clears all stored orders and resets id.
func clearAllStockOrders() {
	stockOrderId = 0
	ordersByCode = make(map[uuid.UUID]domain.StockOrderEntity)
	ordersByCustomer = make(map[uint][]domain.StockOrderEntity)
}