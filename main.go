package main

import (
	"net/http"

	"github.com/eaneto/stocker/handler"
)

var stockHandler handler.Handler
var customerHandler handler.Handler
var stockOrderHandler handler.Handler

// init This is probably the most rudimentary way of doing dependency
// injection.
func init() {
	stockHandler = handler.NewStockHandler()
	customerHandler = handler.CustomerHandler{}
	stockOrderHandler = handler.StockOrderHandler{}
}

func main() {
	http.HandleFunc("/stocks", stockHandler.Handle)
	http.HandleFunc("/stocks/", stockHandler.Handle)
	http.HandleFunc("/customers", customerHandler.Handle)
	http.HandleFunc("/orders", stockOrderHandler.Handle)
	http.ListenAndServe(":8888", nil)
}
