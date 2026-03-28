package service

import (
	"testing"

	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type mockProductRepository struct {
	createFunc   func(name string, price float64) domain.Product
	listFunc     func() []domain.Product
	findByIDFunc func(id string) (domain.Product, bool)
}

func (m *mockProductRepository) Create(name string, price float64) domain.Product {
	return m.createFunc(name, price)
}

func (m *mockProductRepository) List() []domain.Product {
	return m.listFunc()
}

func (m *mockProductRepository) FindByID(id string) (domain.Product, bool) {
	return m.findByIDFunc(id)
}

func TestProductService_Create(t *testing.T) {
	t.Run("should create product successfuly when data is valid", func(t *testing.T) {
		// Setup mock
		repository := &mockProductRepository{
			createFunc: func(name string, price float64) domain.Product {
				return domain.Product{ID: "1", Name: name, Price: price}
			},
		}

		service := NewProductService(repository)

		// Action
		product, err := service.Create("  Monitor 4k  ", 1500.00)
		if err != nil {
			t.Errorf("expected error, got %v", err)
		}
		if product.Name != "Monitor 4k" {
			t.Errorf("expected 'Monitor 4k', got %s", product.Name)
		}
	})

	t.Run("should return fail when name is missing", func(t *testing.T) {
		service := NewProductService(&mockProductRepository{})

		// Action
		_, err := service.Create("", 123.45)
		if err == nil || err.Error() != "Product name is missing" {
			t.Errorf("expected 'Product name is mising' error, got %v", err)
		}
	})

	t.Run("should return fail when price is less than or equal to 0", func(t *testing.T) {
		service := NewProductService(&mockProductRepository{})

		// Action
		_, err := service.Create("Monitor 4k", 0)
		if err == nil || err.Error() != "Price must be greater than zero" {
			t.Errorf("expected 'Product name is mising' error, got %v", err)
		}

		_, err = service.Create("Monitor 4k", -1)
		if err == nil || err.Error() != "Price must be greater than zero" {
			t.Errorf("expected 'Product name is mising' error, got %v", err)
		}
	})
}

func TestProductService_FindByID(t *testing.T) {
	mockedRepo := &mockProductRepository{
		findByIDFunc: func(id string) (domain.Product, bool) {
			products := map[string]domain.Product{
				"1": {ID: "1", Name: "Monitor 4k", Price: 12.0},
			}
			if p, ok := products[id]; ok {
				return p, true
			}
			return domain.Product{}, false
		},
	}
	t.Run("should find product by id", func(t *testing.T) {

		service := NewProductService(mockedRepo)

		// Action
		p, ok := service.FindByID("1")
		if !ok || p.ID != "1" {
			t.Error("expected to find product with ID 1")
		}
	})

	t.Run("should not find product by id when it doesnt exist", func(t *testing.T) {

		service := NewProductService(mockedRepo)

		// Action
		_, ok := service.FindByID("2")
		if ok {
			t.Error("expected not to find product ID 2")
		}
	})
}
