package httphandler

import (
	"go.uber.org/zap"
	"net/http"
)

// Handler for http requests
type Handler struct {
	mux *http.ServeMux
}

// New http handler
func New(s *http.ServeMux, sLogger *zap.SugaredLogger) *Handler {
	h := Handler{s}
	h.registerRoutes()

	return &h
}

// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)
}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}
