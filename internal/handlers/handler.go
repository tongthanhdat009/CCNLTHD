package handlers

import (
    "net/http"
)

// Handler is a struct that holds dependencies for the handler functions.
type Handler struct{}

// NewHandler creates a new Handler instance.
func NewHandler() *Handler {
    return &Handler{}
}

// ExampleHandler is an example HTTP handler function.
func (h *Handler) ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}