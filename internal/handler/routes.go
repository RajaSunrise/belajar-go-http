package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Routes sets up the routing for the Todo API.
func (h *TodoHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", h.ListTodos)
		r.Post("/", h.CreateTodo)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetTodo)
			r.Put("/", h.UpdateTodo)
			r.Delete("/", h.DeleteTodo)
		})
	})

	return r
}
