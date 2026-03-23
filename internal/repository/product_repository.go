package repository

import (
	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type ProductRepository interface {
	List() []domain.Product
	Create(name string, price float64) domain.Product
	FindByID(id string) (domain.Product, bool)
}
