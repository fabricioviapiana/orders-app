package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/service"
)

type orderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *orderHandler {
	return &orderHandler{
		service: service,
	}
}

type createOrderInput struct {
	UserID   string           `json:"user_id"`
	Products []domain.Product `json:"products"`
}

func (h *orderHandler) HandleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func (h *orderHandler) list(w http.ResponseWriter, r *http.Request) {
	orders := h.service.List()
	respondWithJSON(w, http.StatusOK, orders)
}

func (h *orderHandler) create(w http.ResponseWriter, r *http.Request) {
	var createOrderInput createOrderInput

	if err := json.NewDecoder(r.Body).Decode(&createOrderInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	newOrder, err := h.service.Create(createOrderInput.UserID, createOrderInput.Products)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newOrder)
}
