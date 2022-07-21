// Package server is a convenient wrapper over ListenAndServe
package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	// addr - contains the server address
	addr string
	// handler - composable HTTP services with a large set of handlers.
	handler *chi.Mux
	// s- defines parameters for running an HTTP server.
	s *http.Server
}

// New is the server constructor
func New(addr string, router *chi.Mux) *Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &Server{
		addr:    addr,
		handler: router,
		s:       srv,
	}
}

// Start is the method to start the server
func (s *Server) Start() error {
	err := s.s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Shutdown is the method to stop the server
func (s *Server) Shutdown(ctx context.Context) error {
	err := s.s.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
