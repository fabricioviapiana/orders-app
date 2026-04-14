package service

import (
	"errors"
	"strings"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) List() ([]domain.Product, error) {
	return s.repo.List()
}

func (s *ProductService) Create(name string, price float64) (domain.Product, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return domain.Product{}, errors.New("Product name is missing")
	}

	if price <= 0 {
		return domain.Product{}, errors.New("Price must be greater than zero")
	}

	return s.repo.Create(name, price)
}

func (s *ProductService) FindByID(id string) (domain.Product, error) {
	return s.repo.FindByID(id)
}
