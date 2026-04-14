package repository

import "github.com/fabricioviapiana/orders-app/internal/domain"

type UserRepository interface {
	Create(name, email string) (domain.User, error)
	List() ([]domain.User, error)
	FindByID(id string) (domain.User, error)
}
