package todos

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type jsonError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"msg"`
}

type Handler struct {
	DB Storage
}

func (h Handler) GetCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		todos := h.DB.GetAll()
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var t struct {
			Task string `json:"task"`
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo := h.DB.Create(t.Task)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	default:
		errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h Handler) GetUpdateDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// r.URL.Path == /api/todos/:id
	p := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(p) != 3 {
		errorResponse(w, "bad url", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p[2])
	if err != nil {
		errorResponse(w, "bad id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		todo, err := h.DB.Get(id)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(todo)
	case http.MethodPut:
		var t struct {
			Task string `json:"task"`
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo, err := h.DB.Update(id, t.Task)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(todo)
	case http.MethodDelete:
		todo, err := h.DB.Delete(id)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(todo)
	default:
		errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func errorResponse(w http.ResponseWriter, msg string, sc int) {
	err := jsonError{
		StatusCode: sc,
		Message:    msg,
	}
	w.WriteHeader(sc)
	json.NewEncoder(w).Encode(err)
}
