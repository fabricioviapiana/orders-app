package service

import (
	"errors"
	"strings"

	"github.com/fabricioviapiana/orders-app/internal/domain"
	"github.com/fabricioviapiana/orders-app/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(name, email string) (*domain.User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return nil, errors.New("User name is missing")
	}

	if email == "" {
		return nil, errors.New("Email is missing")
	}

	newUser := s.repo.Create(name, email)

	return &newUser, nil
}

func (s *UserService) List() []domain.User {
	return s.repo.List()
}

func (s *UserService) FindByID(id string) (domain.User, bool) {
	return s.repo.FindByID(id)
}
