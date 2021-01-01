package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eaneto/stocker/controller"
	"github.com/eaneto/stocker/domain"
	"github.com/sirupsen/logrus"
)

type StockHandler struct {
	StockController controller.BaseStockController
}

func NewStockHandler() Handler {
	return StockHandler{
		StockController: controller.NewStockController(),
	}
}

func (handler StockHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handler.handleGet(w, r)
	} else if r.Method == http.MethodPost {
		handler.handlePost(w, r)
	} else if r.Method == http.MethodPatch {
		logrus.Info("PATCH stock")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler StockHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	ticker := strings.TrimPrefix(r.URL.Path, "/stocks")
	if ticker == "" || ticker == "/" {
		logrus.Info("GET all stocks")
		handler.getAllStocks(w, r)
	} else {
		ticker = strings.ReplaceAll(ticker, "/", "")
		handler.getSpecificStock(ticker, w, r)
	}
}

func (handler StockHandler) getAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, status := handler.StockController.FindAll()
	payload, err := json.Marshal(stocks)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Marshaling stocks to JSON.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(status)
}

func (handler StockHandler) getSpecificStock(ticker string, w http.ResponseWriter, r *http.Request) {
	logrus.WithField("ticker", ticker).Info("GET specific stock")
	stock, status := handler.StockController.FindByTicker(ticker)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	} else {
		payload, err := json.Marshal(stock)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("Marshaling stock.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logrus.WithField("payload", string(payload)).
			Info("Found stock.")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (handler StockHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	logrus.Info("POST stock")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Reading request body.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stock := domain.Stock{}
	err = json.Unmarshal(body, &stock)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Unmarshaling JSON.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := handler.StockController.RegisterStock(stock)
	w.WriteHeader(status)
}
