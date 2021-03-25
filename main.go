package main

import (
	"net/http"

	"github.com/eaneto/stocker/customer"
	"github.com/eaneto/stocker/order"
	"github.com/eaneto/stocker/stock"
)

var stockHandler http.Handler
var customerHandler http.Handler
var stockOrderHandler http.Handler
var customerPositionHandler http.Handler

func init() {
	stockHandler = stock.NewStockHandler()
	customerHandler = customer.NewCustomerHandler()
	stockOrderHandler = order.NewStockOrderHandler()
	customerPositionHandler = order.NewCustomerPositionHandler()
}

func main() {
	stockOrderService := order.NewStockOrderService()
	go stockOrderService.ConfirmOrders()

	http.HandleFunc("/stocks", stockHandler.ServeHTTP)
	http.HandleFunc("/stocks/", stockHandler.ServeHTTP)
	http.HandleFunc("/customers", customerHandler.ServeHTTP)
	http.HandleFunc("/orders", stockOrderHandler.ServeHTTP)
	http.HandleFunc("/orders/", stockOrderHandler.ServeHTTP)
	http.HandleFunc("/positions/", customerPositionHandler.ServeHTTP)
	http.ListenAndServe(":8888", nil)
}
