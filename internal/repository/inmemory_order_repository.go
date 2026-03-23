package repository

import (
	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var orders = []domain.Order{}

type inMemoryOrderRepository struct{}

func NewInMemoryOrderRepository() *inMemoryOrderRepository {
	return &inMemoryOrderRepository{}
}

func (r *inMemoryOrderRepository) Create(userID string, items []domain.OrderItem, totalAmount float64) domain.Order {
	newOrder := domain.Order{
		ID:          uuid.NewString(),
		UserID:      userID,
		Items:       items,
		TotalAmount: totalAmount,
	}

	orders = append(orders, newOrder)

	return newOrder
}

func (r *inMemoryOrderRepository) List() []domain.Order {
	return orders
}
