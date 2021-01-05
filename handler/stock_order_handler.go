package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
	customerID := strings.TrimPrefix(r.URL.Path, "/orders")
	if customerID == "" || customerID == "/" {
		logrus.Info("GET all orders")
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
