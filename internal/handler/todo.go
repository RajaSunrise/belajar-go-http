package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"belajar-go-http/internal/model"
)

// TodoHandler holds the dependencies for the todo handlers.
type TodoHandler struct {
	Store *model.TodoStore
}

// respondWithError is a helper function for sending uniform error responses.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON is a helper function for sending JSON responses.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ListTodos handles GET /todos requests.
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.Store.List()
	respondWithJSON(w, http.StatusOK, todos)
}

// CreateTodo handles POST /todos requests.
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Task string `json:"task"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if request.Task == "" {
		respondWithError(w, http.StatusBadRequest, "Task cannot be empty")
		return
	}

	todo := h.Store.Create(request.Task)
	respondWithJSON(w, http.StatusCreated, todo)
}

// GetTodo handles GET /todos/{id} requests.
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	todo, found := h.Store.Get(id)
	if !found {
		respondWithError(w, http.StatusNotFound, "Todo not found")
		return
	}

	respondWithJSON(w, http.StatusOK, todo)
}

// UpdateTodo handles PUT /todos/{id} requests.
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var request struct {
		Task      string `json:"task"`
		Completed bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if request.Task == "" {
		respondWithError(w, http.StatusBadRequest, "Task cannot be empty")
		return
	}

	todo, found := h.Store.Update(id, request.Task, request.Completed)
	if !found {
		respondWithError(w, http.StatusNotFound, "Todo not found")
		return
	}

	respondWithJSON(w, http.StatusOK, todo)
}

// DeleteTodo handles DELETE /todos/{id} requests.
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	if err := h.Store.Delete(id); err != nil {
		respondWithError(w, http.StatusNotFound, "Todo not found")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
