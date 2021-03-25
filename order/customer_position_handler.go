package order

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomerPositionHandler struct {
	StockOrderService BaseStockOrderService
}

func NewCustomerPositionHandler() CustomerPositionHandler {
	return CustomerPositionHandler{
		StockOrderService: NewStockOrderService(),
	}
}

func (handler CustomerPositionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler.handleGet(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler CustomerPositionHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	customerID := strings.TrimPrefix(r.URL.Path, "/positions")
	if customerID == "" || customerID == "/" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		logrus.Info("GET customer position")
		customerID = strings.ReplaceAll(customerID, "/", "")
		customer, _ := strconv.Atoi(customerID)
		position, _ := handler.StockOrderService.GetCustomerPosition(uint(customer))
		payload, _ := json.Marshal(position)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
