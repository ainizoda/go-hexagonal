package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ainizoda/go-hexagonal/internal/adapters/in/http/dto"
	"github.com/ainizoda/go-hexagonal/internal/domain/user"
)

type UserHandler struct {
	svc user.InputPort
}

func NewUserHandler(svc user.InputPort) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", h.List)
	mux.HandleFunc("GET /api/user", h.Get)
	mux.HandleFunc("POST /api/user", h.Add)
	mux.HandleFunc("DELETE /api/user", h.Delete)
}

func (uc *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	if id == "" {
		http.Error(w, "'id' param is required", http.StatusBadRequest)
		return
	}
	usr, err := uc.svc.Get(context.Background(), id)
	if err != nil {
		if errors.Is(err, user.ErrUserDoesNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(usr)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (uc *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var data dto.UserRequestBody
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	usr, err := user.New(data.FirstName, data.LastName, data.Email, data.Roles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := uc.svc.Add(context.Background(), usr); err != nil {
		if errors.Is(err, user.ErrUserAlreadyExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	if id == "" {
		http.Error(w, "'id' param is required", http.StatusBadRequest)
		return
	}
	if err := uc.svc.Delete(context.Background(), id); err != nil {
		if errors.Is(err, user.ErrUserDoesNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := uc.svc.List(context.Background())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&dto.UserResponseBody{Data: users})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
