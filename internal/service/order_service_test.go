package service

import (
	"testing"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/repository"
)

// Mock para o UserService usando o ISP
type mockUserService struct {
	findFunc func(id string) (domain.User, bool)
}

func (m *mockUserService) FindByID(id string) (domain.User, bool) {
	return m.findFunc(id)
}

func TestOrderService_Create(t *testing.T) {
	// Setup
	orderRepo := repository.NewInMemoryOrderRepository()
	productRepo := repository.NewInMemoryProductRepository()
	productSvc := NewProductService(productRepo)

	// Mock do UserService
	userSvc := &mockUserService{
		findFunc: func(id string) (domain.User, bool) {
			users := map[string]domain.User{
				"user-123": {ID: "user-123"},
			}
			user, ok := users[id]
			if ok {
				return user, true
			}
			return domain.User{}, false
		},
	}

	// Criando um produto para o teste
	p1 := productRepo.Create("Teclado Mecânico", 250.0)
	p2 := productRepo.Create("Mouse Gamer", 150.0)

	// Agora passamos os 3 argumentos
	s := NewOrderService(orderRepo, productSvc, userSvc)

	t.Run("should create an order successfully and calculate total amount correctly", func(t *testing.T) {
		input := CreateOrderInput{
			UserID: "user-123",
			Items: []CreateOrderItemInput{
				{ProductID: p1.ID, Quantity: 2}, // 500.0
				{ProductID: p2.ID, Quantity: 1}, // 150.0
			},
		}

		order, err := s.Create(input)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expectedTotal := 650.0
		if order.TotalAmount != expectedTotal {
			t.Errorf("expected total amount %.2f, got %.2f", expectedTotal, order.TotalAmount)
		}

		if len(order.Items) != 2 {
			t.Errorf("expected 2 items, got %d", len(order.Items))
		}
	})

	t.Run("should return error when user is not found", func(t *testing.T) {
		input := CreateOrderInput{
			UserID: "user-999", // User que o nosso mock vai dizer que não existe
			Items: []CreateOrderItemInput{
				{ProductID: p1.ID, Quantity: 1},
			},
		}

		_, err := s.Create(input)

		if err == nil {
			t.Error("expected an error, got nil")
		}

		expectedError := "user user-999 not found"
		if err.Error() != expectedError {
			t.Errorf("expected error %q, got %q", expectedError, err.Error())
		}
	})

	t.Run("should return error when product is not found", func(t *testing.T) {
		input := CreateOrderInput{
			UserID: "user-123",
			Items: []CreateOrderItemInput{
				{ProductID: "non-existent-id", Quantity: 1},
			},
		}

		_, err := s.Create(input)

		if err == nil {
			t.Error("expected an error, got nil")
		}

		expectedError := "product non-existent-id not found"
		if err.Error() != expectedError {
			t.Errorf("expected error %q, got %q", expectedError, err.Error())
		}
	})

	t.Run("should return error when quantity is less or equal 0", func(t *testing.T) {
		input := CreateOrderInput{
			UserID: "user-123",
			Items: []CreateOrderItemInput{
				{ProductID: p1.ID, Quantity: 0},
			},
		}

		_, err := s.Create(input)
		if err == nil {
			t.Errorf("expected error but got nil")
		}

		expectedError := "products must have quantity greather than 0"
		if err.Error() != expectedError {
			t.Errorf("expected error %q, got %q", expectedError, err.Error())
		}
	})
}
