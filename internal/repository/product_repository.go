package repository

import (
	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type ProductRepository interface {
	List() ([]domain.Product, error)
	Create(name string, price float64) (domain.Product, error)
	FindByID(id string) (domain.Product, error)
}
