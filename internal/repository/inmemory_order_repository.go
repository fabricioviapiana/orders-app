package repository

import (
	"fmt"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var orders = []domain.Order{}

type inMemoryOrderRepository struct{}

func NewInMemoryOrderRepository() *inMemoryOrderRepository {
	return &inMemoryOrderRepository{}
}

func (r *inMemoryOrderRepository) Create(userID string, items []domain.OrderItem, totalAmount float64) (domain.Order, error) {
	newOrder := domain.Order{
		ID:          uuid.NewString(),
		UserID:      userID,
		Items:       items,
		TotalAmount: totalAmount,
	}

	orders = append(orders, newOrder)

	return newOrder, nil
}

func (r *inMemoryOrderRepository) List() ([]domain.Order, error) {
	return orders, nil
}

func (r *inMemoryOrderRepository) FindByID(id string) (domain.Order, error) {
	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}
	return domain.Order{}, fmt.Errorf("order not found")
}
