package service

import (
	"errors"
	"fmt"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/repository"
)

type productService interface {
	FindByID(id string) (domain.Product, error)
}

type userService interface {
	FindByID(id string) (domain.User, error)
}

type OrderService struct {
	orderRepository repository.OrderRepository
	productService  productService
	userService     userService
}

func NewOrderService(orderRepository repository.OrderRepository, productService productService, userService userService) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productService:  productService,
		userService:     userService,
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

	_, err := s.userService.FindByID(input.UserID)
	if err != nil {
		return nil, fmt.Errorf("user %s not found", input.UserID)
	}

	var totalAmount float64
	var orderItems []domain.OrderItem

	for _, item := range input.Items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("products must have quantity greather than 0")
		}

		product, err := s.productService.FindByID(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product %s not found", item.ProductID)
		}

		orderItems = append(orderItems, domain.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			UnitPrice: product.Price,
		})
		totalAmount += product.Price * float64(item.Quantity)
	}

	newOrder, err := s.orderRepository.Create(input.UserID, orderItems, totalAmount)
	if err != nil {
		return nil, err
	}

	return &newOrder, nil
}

func (s *OrderService) List() ([]domain.Order, error) {
	return s.orderRepository.List()
}

func (s *OrderService) FindByID(id string) (domain.Order, error) {
	return s.orderRepository.FindByID(id)
}
