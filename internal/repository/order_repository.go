package repository

import "github.com/fabricioviapiana/orders-app/internal/domain"

type OrderRepository interface {
	Create(userID string, products []domain.Product, totalAmount float64) domain.Order
	List() []domain.Order
}
