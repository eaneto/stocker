package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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
		handler.handleGet(w, r)
	} else if r.Method == "POST" {
		handler.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler StockOrderHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	orderCode := strings.TrimPrefix(r.URL.Path, "/orders")
	if orderCode == "" || orderCode == "/" {
		logrus.Info("GET all orders")
	} else {
		orderCode = strings.ReplaceAll(orderCode, "/", "")
		logrus.Info("GET order info " + orderCode)
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
	_, alreadyProcessed := err.(domain.StockOrderAlreadyProcessedError)
	if alreadyProcessed {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Creating order.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
