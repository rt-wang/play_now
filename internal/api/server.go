package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Server holds the HTTP server configuration.
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new API server with default routes.
func NewServer() *Server {
	r := mux.NewRouter()
	// placeholder routes
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	srv := &http.Server{Addr: ":8080", Handler: r}
	return &Server{httpServer: srv}
}

// Run starts the HTTP server.
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
