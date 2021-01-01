package main

import (
	"net/http"

	"github.com/eaneto/stocker/controller"
	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/handler"
	"github.com/eaneto/stocker/repository"
	"github.com/eaneto/stocker/service"
	"github.com/eaneto/stocker/usecase"
)

var stockHandler handler.Handler
var customerHandler handler.Handler
var stockOrderHandler handler.Handler

// init This is probably the most rudimentary way of doing dependency
// injection.
func init() {
	stockHandler = handler.StockHandler{
		StockController: controller.StockController{
			StockUseCase: usecase.StockUseCase{
				StockService: service.StockService{
					StockRepository: repository.StockRepository{
						Stocks: make(map[string]domain.StockEntity),
					},
				},
			},
		},
	}
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
