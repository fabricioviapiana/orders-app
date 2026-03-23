package service

import (
	"errors"
	"fmt"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/repository"
)

type productService interface {
	FindByID(id string) (domain.Product, bool)
}

type OrderService struct {
	orderRepository repository.OrderRepository
	productService  productService
}

func NewOrderService(orderRepository repository.OrderRepository, productService productService) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productService:  productService,
	}
}

type CreateOrderItemInput struct {
	ProductID string
	Quantity  int
}

type CreateOrderInput struct {
	UserID string
	Items  []CreateOrderItemInput
}

func (s *OrderService) Create(input CreateOrderInput) (*domain.Order, error) {
	if input.UserID == "" {
		return nil, errors.New("User ID is missing")
	}

	if len(input.Items) == 0 {
		return nil, errors.New("Order must have at least one product")
	}

	var totalAmount float64
	var orderItems []domain.OrderItem

	for _, item := range input.Items {
		product, ok := s.productService.FindByID(item.ProductID)
		if !ok {
			return nil, fmt.Errorf("product %s not found", item.ProductID)
		}
		orderItems = append(orderItems, domain.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			UnitPrice: product.Price,
		})
		totalAmount += product.Price * float64(item.Quantity)
	}

	newOrder := s.orderRepository.Create(input.UserID, orderItems, totalAmount)

	return &newOrder, nil
}

func (s *OrderService) List() []domain.Order {
	return s.orderRepository.List()
}
