package domain

import (
	"context"

	"gorm.io/gorm"
)

// User represents the database entity model for a user
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
}

// UserRepository defines the database operations contract for User
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id uint) (*User, error)
}

// UserService defines the business logic operations contract for User
type UserService interface {
	Register(ctx context.Context, name, email string) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
}
