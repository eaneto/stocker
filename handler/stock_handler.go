package handler

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type StockHandler struct{}

func (StockHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ticker := strings.TrimPrefix(r.URL.Path, "/stocks/")
	logrus.WithField("ticker", ticker).Info("ticker")

	if r.Method == "GET" {
		logrus.Info("GET stock")
	} else if r.Method == "POST" {
		logrus.Info("POST stock")
	} else if r.Method == "PATCH" {
		logrus.Info("PATCH stock")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
