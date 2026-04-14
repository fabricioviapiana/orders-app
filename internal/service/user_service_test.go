package service

import (
	"testing"

	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type mockUserRepository struct {
	createFunc func(name, email string) (domain.User, error)
	listFunc   func() ([]domain.User, error)
	findFunc   func(id string) (domain.User, error)
}

func (m *mockUserRepository) Create(name, email string) (domain.User, error) {
	return m.createFunc(name, email)
}

func (m *mockUserRepository) List() ([]domain.User, error) {
	return m.listFunc()
}

func (m *mockUserRepository) FindByID(id string) (domain.User, error) {
	return m.findFunc(id)
}

func TestUserService_Create(t *testing.T) {
	t.Run("should create a user successfuly when data is valid", func(t *testing.T) {
		userRepoMock := &mockUserRepository{
			createFunc: func(name, email string) (domain.User, error) {
				return domain.User{
					ID:    "abc",
					Name:  name,
					Email: email,
				}, nil
			},
		}

		userSvc := NewUserService(userRepoMock)

		_, err := userSvc.Create("John", "john@email.com")
		if err != nil {
			t.Errorf("not expected error, got %v", err)
		}
	})

	t.Run("should fail if name wasn't provided", func(t *testing.T) {
		userSvc := NewUserService(&mockUserRepository{})

		_, err := userSvc.Create("", "email")
		if err == nil || err.Error() != "User name is missing" {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("should fail if email wasn't provided", func(t *testing.T) {
		userSvc := NewUserService(&mockUserRepository{})

		_, err := userSvc.Create("Hoop", "")
		if err == nil || err.Error() != "Email is missing" {
			t.Errorf("expected error, got %v", err)
		}
	})
}
