package repository

import (
	"fmt"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var products = []domain.Product{}

type inMemoryProductRepository struct{}

func NewInMemoryProductRepository() *inMemoryProductRepository {
	return &inMemoryProductRepository{}
}

func (r *inMemoryProductRepository) List() ([]domain.Product, error) {
	return products, nil
}

func (r *inMemoryProductRepository) Create(name string, price float64) (domain.Product, error) {
	newProduct := domain.Product{ID: uuid.NewString(), Name: name, Price: price}
	products = append(products, newProduct)

	return newProduct, nil
}

func (r *inMemoryProductRepository) FindByID(id string) (domain.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return domain.Product{}, fmt.Errorf("product not found")
}
