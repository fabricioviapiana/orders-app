package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fabricioviapiana/orders-app/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

type createProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		if id := r.PathValue("id"); id != "" {
			h.find(w, r)
			return
		}
		h.list(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) list(w http.ResponseWriter, r *http.Request) {
	products := h.service.List()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "error listing products", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	var input createProductInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	product, err := h.service.Create(input.Name, input.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) find(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "missing product id")
		return
	}

	product, ok := h.service.FindByID(id)
	if !ok {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	respondWithJSON(w, http.StatusOK, product)
}
