package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eaneto/stocker/domain"
	"github.com/eaneto/stocker/service"
	"github.com/sirupsen/logrus"
)

type StockOrderHandler struct {
	StockOrderService service.BaseStockOrderService
}

func NewStockOrderHandler() http.Handler {
	return StockOrderHandler{
		StockOrderService: service.NewStockOrderService(),
	}
}

func (handler StockOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logrus.Info("GET order")
	} else if r.Method == "POST" {
		handler.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler StockOrderHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	logrus.Info("POST order")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Reading request body.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stockOrderRequest := domain.StockOrderRequest{}
	err = json.Unmarshal(body, &stockOrderRequest)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Unmarshaling JSON.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = handler.StockOrderService.CreateOrder(stockOrderRequest)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Creating order.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
