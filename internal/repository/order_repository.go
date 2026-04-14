package repository

import "github.com/fabricioviapiana/orders-app/internal/domain"

type OrderRepository interface {
	Create(userID string, items []domain.OrderItem, totalAmount float64) (domain.Order, error)
	List() ([]domain.Order, error)
	FindByID(id string) (domain.Order, error)
}
