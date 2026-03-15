package repository

import (
	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var products = []domain.Product{}

type InMemoryProductRepository struct{}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{}
}

func (r InMemoryProductRepository) List() []domain.Product {
	return products
}

func (r *InMemoryProductRepository) Create(name string, price float64) domain.Product {
	newProduct := domain.Product{ID: uuid.NewString(), Name: name, Price: price}
	products = append(products, newProduct)

	return newProduct
}
