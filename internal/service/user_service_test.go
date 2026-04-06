package service

import (
	"testing"

	"github.com/fabricioviapiana/orders-app/internal/domain"
)

type mockUserRepository struct {
	createFunc func(name, email string) domain.User
	listFunc   func() []domain.User
	findFunc   func(id string) (domain.User, bool)
}

func (m *mockUserRepository) Create(name, email string) domain.User {
	return m.createFunc(name, email)
}

func (m *mockUserRepository) List() []domain.User {
	return m.listFunc()
}

func (m *mockUserRepository) FindByID(id string) (domain.User, bool) {
	return m.findFunc(id)
}

func TestUserService_Create(t *testing.T) {
	t.Run("should create a user successfuly when data is valid", func(t *testing.T) {
		userRepoMock := &mockUserRepository{
			createFunc: func(name, email string) domain.User {
				return domain.User{
					ID:    "abc",
					Name:  name,
					Email: email,
				}
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
