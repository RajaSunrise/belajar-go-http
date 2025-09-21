package handler

import (
	"html/template"
	"log"
	"net/http"

	"belajar-go-http/internal/model"
)

// WebHandler serves the HTML web interface.
func (h *TodoHandler) WebHandler(w http.ResponseWriter, r *http.Request) {
	// Path to the template file
	// Note: This path is relative to the project root where the binary is run.
	tmplPath := "ui/html/todos.html"

	// Parse the template. It's good practice to parse templates once at startup,
	// but for this example, we'll parse it on each request for simplicity.
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("Error parsing template %s: %v", tmplPath, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get all the todos from the store
	todos := h.Store.List()

	// Create a data structure to pass to the template
	data := struct {
		Todos []model.Todo
	}{
		Todos: todos,
	}

	// Execute the template with the data and write the output to the ResponseWriter
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
