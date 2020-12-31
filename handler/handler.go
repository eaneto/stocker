package handler

import (
	"net/http"
)

// Handler Basic interface for HTTP handlers.
type Handler interface {
	// Handle Handles a HTTP Request.
	Handle(w http.ResponseWriter, r *http.Request)
}
