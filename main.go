package main

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func stockHandler(w http.ResponseWriter, r *http.Request) {
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

func customerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logrus.Info("GET customer")
	} else if r.Method == "POST" {
		logrus.Info("POST customer")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		logrus.Info("GET order")
	} else if r.Method == "POST" {
		logrus.Info("POST order")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/stocks", stockHandler)
	http.HandleFunc("/stocks/", stockHandler)
	http.HandleFunc("/customers", customerHandler)
	http.HandleFunc("/orders", orderHandler)
	http.ListenAndServe(":8888", nil)
}
