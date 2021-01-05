package main

import (
	"net/http"

	"github.com/eaneto/stocker/handler"
	"github.com/eaneto/stocker/service"
)

var stockHandler http.Handler
var customerHandler http.Handler
var stockOrderHandler http.Handler
var customerPositionHandler http.Handler

func init() {
	stockHandler = handler.NewStockHandler()
	customerHandler = handler.NewCustomerHandler()
	stockOrderHandler = handler.NewStockOrderHandler()
	customerPositionHandler = handler.NewCustomerPositionHandler()
}

func main() {
	stockOrderService := service.NewStockOrderService()
	go stockOrderService.ConfirmOrders()

	http.HandleFunc("/stocks", stockHandler.ServeHTTP)
	http.HandleFunc("/stocks/", stockHandler.ServeHTTP)
	http.HandleFunc("/customers", customerHandler.ServeHTTP)
	http.HandleFunc("/orders", stockOrderHandler.ServeHTTP)
	http.HandleFunc("/orders/", stockOrderHandler.ServeHTTP)
	http.HandleFunc("/positions/", customerPositionHandler.ServeHTTP)
	http.ListenAndServe(":8888", nil)
}
