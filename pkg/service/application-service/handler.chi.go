package applicationservice

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewHandler returns an http Handler with initilized application routes.
func NewHandler(service Service) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Post("/api/applications", service.CreateApplication)
	r.Get("/api/applications", service.ListApplication)

	return r
}
