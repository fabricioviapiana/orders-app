package repository

import "github.com/fabricioviapiana/orders-app/internal/domain"

type UserRepository interface {
	Create(name, email string) domain.User
	List() []domain.User
}
