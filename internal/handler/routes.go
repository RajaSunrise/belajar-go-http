package handler

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Routes sets up the routing for the Todo application.
func (h *TodoHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// --- Web Interface Routes ---
	r.Get("/", h.WebHandler)

	// --- Static File Server ---
	// Create a file server handler to serve static files
	staticDir, _ := filepath.Abs("./ui/static")
	fileServer(r, "/static", http.Dir(staticDir))

	// --- API Routes ---
	r.Route("/api/todos", func(r chi.Router) {
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

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a directory.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("fileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
