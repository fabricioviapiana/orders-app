package repository

import (
	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/google/uuid"
)

var users = []domain.User{}

type inMemoryUserRepository struct{}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	return &inMemoryUserRepository{}
}

func (r *inMemoryUserRepository) List() []domain.User {
	return users
}

func (r *inMemoryUserRepository) Create(name, email string) domain.User {
	newUser := domain.User{ID: uuid.NewString(), Name: name, Email: email}
	users = append(users, newUser)

	return newUser
}

func (r *inMemoryUserRepository) FindByID(id string) (domain.User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}

	return domain.User{}, false
}
