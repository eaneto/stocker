package customer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	CustomerService BaseCustomerService
}

func NewCustomerHandler() http.Handler {
	return CustomerHandler{
		CustomerService: NewCustomerService(),
	}
}

func (handler CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler.handleGet(w, r)
	} else if r.Method == "POST" {
		handler.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler CustomerHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	logrus.Info("GET customers")
	customers := handler.CustomerService.FindAll()
	payload, err := json.Marshal(customers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Marshaling customers to JSON.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (handler CustomerHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	logrus.Info("POST customer")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Reading request body.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	customer := Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Unmarshaling JSON.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.CustomerService.Create(customer.Name)
	w.WriteHeader(http.StatusCreated)
}
