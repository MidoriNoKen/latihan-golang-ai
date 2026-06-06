package service

import (
	"context"
	"errors"

	"github.com/MidoriNoKen/latihan-golang-ai/domain"
)

type userService struct {
	repo domain.UserRepository
}

// NewUserService creates a new UserService implementation with repository dependency
func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

// Register contains logic to create a new user with validation
func (s *userService) Register(ctx context.Context, name, email string) (*domain.User, error) {
	if name == "" {
		return nil, errors.New("name is required and cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email is required and cannot be empty")
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all registered users
func (s *userService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx)
}

// GetUserByID retrieves a user by their specific ID
func (s *userService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}
