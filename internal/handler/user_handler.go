package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fabricioviapiana/orders-app/internal/service"
)

type userHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *userHandler {
	return &userHandler{
		service: service,
	}
}

type createUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *userHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) list(w http.ResponseWriter, r *http.Request) {
	users := h.service.List()
	respondWithJSON(w, http.StatusOK, users)
}

func (h *userHandler) create(w http.ResponseWriter, r *http.Request) {
	var createUserInput createUserInput

	if err := json.NewDecoder(r.Body).Decode(&createUserInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	newUser, err := h.service.Create(createUserInput.Name, createUserInput.Email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newUser)
}
