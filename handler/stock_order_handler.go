package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type StockOrderHandler struct{}

func NewStockOrderHandler() http.Handler {
	return StockOrderHandler{}
}

func (StockOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logrus.Info("GET order")
	} else if r.Method == "POST" {
		logrus.Info("POST order")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
