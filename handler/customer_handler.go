package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type CustomerHandler struct{}

func (CustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logrus.Info("GET customer")
	} else if r.Method == "POST" {
		logrus.Info("POST customer")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
