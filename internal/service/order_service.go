package service

import (
	"errors"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Create(userID string, products []domain.Product) (*domain.Order, error) {
	if userID == "" {
		return nil, errors.New("User ID is missing")
	}

	if len(products) == 0 {
		return nil, errors.New("Order must have at least one product")
	}

	var totalAmount float64
	for _, p := range products {
		totalAmount += p.Price
	}

	newOrder := s.repo.Create(userID, products, totalAmount)

	return &newOrder, nil
}

func (s *OrderService) List() []domain.Order {
	return s.repo.List()
}
