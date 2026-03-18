package repository

import (
	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var products = []domain.Product{}

type inMemoryProductRepository struct{}

func NewInMemoryProductRepository() *inMemoryProductRepository {
	return &inMemoryProductRepository{}
}

func (r inMemoryProductRepository) List() []domain.Product {
	return products
}

func (r *inMemoryProductRepository) Create(name string, price float64) domain.Product {
	newProduct := domain.Product{ID: uuid.NewString(), Name: name, Price: price}
	products = append(products, newProduct)

	return newProduct
}
