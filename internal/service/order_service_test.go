package service

import (
	"testing"

	"github.com/fabricioviapiana/orders-app/internal/repository"
)

func TestOrderService_Create(t *testing.T) {
	// Setup
	orderRepo := repository.NewInMemoryOrderRepository()
	productRepo := repository.NewInMemoryProductRepository()

	// Criando um produto para o teste
	p1 := productRepo.Create("Teclado Mecânico", 250.0)
	p2 := productRepo.Create("Mouse Gamer", 150.0)

	s := NewOrderService(orderRepo, productRepo)

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

	t.Run("should return error when user id is missing", func(t *testing.T) {
		input := CreateOrderInput{
			UserID: "",
			Items: []CreateOrderItemInput{
				{ProductID: p1.ID, Quantity: 1},
			},
		}

		_, err := s.Create(input)

		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}
