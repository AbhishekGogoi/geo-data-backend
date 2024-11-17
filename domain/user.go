package domain

import (
	"context"
)

// User entity represents the domain model for a user.
type User struct {
	ID       string // Use UUID or Serial from PostgreSQL
	Name     string
	Email    string
	Password string
}

// UserRepository defines the interface for user-related operations.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Fetch(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id string) (User, error)
}