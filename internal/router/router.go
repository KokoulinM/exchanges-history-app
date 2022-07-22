// Package router composable HTTP services with a large set of handlers
package router

import (
	"github.com/KokoulinM/exchanges-history-app/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New router constructor
func New(h *handlers.Handlers) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/history/exchanges/{file}", h.UploadHistory)
			r.Get("/history/exchanges", h.GetHistory)
			r.Get("/ping", h.PingDB)
		})
	})

	return router
}
