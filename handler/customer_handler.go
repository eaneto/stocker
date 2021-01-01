package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eaneto/stocker/controller"
	"github.com/eaneto/stocker/domain"
	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	CustomerController controller.BaseCustomerController
}

func NewCustomerHandler() Handler {
	return CustomerHandler{
		CustomerController: controller.NewCustomerController(),
	}
}

func (handler CustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler.handleGet(w, r)
	} else if r.Method == "POST" {
		handler.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler CustomerHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	logrus.Info("GET customer")
	customers, status := handler.CustomerController.FindAll()
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
	w.WriteHeader(status)
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
	customer := domain.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Unmarshaling JSON.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	status := handler.CustomerController.Create(customer.Name)
	w.WriteHeader(status)
}
