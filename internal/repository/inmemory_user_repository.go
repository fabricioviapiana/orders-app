package repository

import (
	"fmt"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var users = []domain.User{}

type inMemoryUserRepository struct{}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	return &inMemoryUserRepository{}
}

func (r *inMemoryUserRepository) List() ([]domain.User, error) {
	return users, nil
}

func (r *inMemoryUserRepository) Create(name, email string) (domain.User, error) {
	newUser := domain.User{ID: uuid.NewString(), Name: name, Email: email}
	users = append(users, newUser)

	return newUser, nil
}

func (r *inMemoryUserRepository) FindByID(id string) (domain.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, fmt.Errorf("user not found")
}
