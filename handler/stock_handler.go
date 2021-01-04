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

type StockHandler struct {
	StockService service.BaseStockService
}

func NewStockHandler() http.Handler {
	return StockHandler{
		StockService: service.NewStockService(),
	}
}

func (handler StockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	stocks, status := handler.findAll()
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
	stock, status := handler.findByTicker(ticker)
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

	status := handler.registerStock(stock)
	w.WriteHeader(status)
}

func (handler StockHandler) registerStock(stock domain.Stock) (httpStatus int) {
	err := handler.StockService.RegisterStock(stock)
	if err == nil {
		return http.StatusCreated
	}
	_, isConflicError := err.(domain.AlreadyRegisteredStockError)
	if isConflicError {
		return http.StatusConflict
	} else {
		return http.StatusInternalServerError
	}
}

func (handler StockHandler) findByTicker(ticker string) (domain.Stock, int) {
	stock, err := handler.StockService.SearchByTicker(ticker)
	if err == nil {
		return domain.Stock{
			Ticker: stock.Ticker,
			Price:  stock.Price,
		}, http.StatusOK
	}
	_, isNotFound := err.(domain.StockNotFoundError)
	if isNotFound {
		return domain.Stock{}, http.StatusNotFound
	}
	return domain.Stock{}, http.StatusInternalServerError
}

func (handler StockHandler) findAll() ([]domain.Stock, int) {
	return handler.StockService.FindAll(), http.StatusOK
}
