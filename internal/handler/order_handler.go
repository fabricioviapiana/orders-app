package handler

import (
	"encoding/json"
	"net/http"

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

type createOrderItemInput struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type createOrderInput struct {
	UserID string                 `json:"userId"`
	Items  []createOrderItemInput `json:"items"`
}

func (h *orderHandler) create(w http.ResponseWriter, r *http.Request) {
	var createOrderInput createOrderInput

	if err := json.NewDecoder(r.Body).Decode(&createOrderInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	serviceInput := service.CreateOrderInput{
		UserID: createOrderInput.UserID,
		Items:  make([]service.CreateOrderItemInput, len(createOrderInput.Items)),
	}

	for i, item := range createOrderInput.Items {
		serviceInput.Items[i] = service.CreateOrderItemInput{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	newOrder, err := h.service.Create(serviceInput)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, newOrder)
}
