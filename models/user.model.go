package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName string    `gorm:"type:varchar(255);not null"`                      // User's first name
	LastName  string    `gorm:"type:varchar(255);not null"`                      // User's last name
	Email     string    `gorm:"uniqueIndex;not null"`                           // Unique email for login
	Password  string    `gorm:"not null"`                                       // Hashed password
	CreatedAt time.Time `gorm:"not null"`                                       // Timestamp for when the user was created
	UpdatedAt time.Time `gorm:"not null"`
}

type SignUpInput struct {
	FirstName       string `json:"first_name" binding:"required"`         // First name is required
	LastName        string `json:"last_name" binding:"required"`          // Last name is required
	Email           string `json:"email" binding:"required,email"`        // Valid email address is required
	Password        string `json:"password" binding:"required,min=8"`     // Password must have a minimum length of 8
	PasswordConfirm string `json:"password_confirm" binding:"required"`   // Confirmation for the password
}


type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`               // Valid email address is required
	Password string `json:"password" binding:"required"`                  // Password is required
}

// UserResponse represents the safe fields returned in API responses.
type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`                            // User ID
	FirstName string    `json:"first_name,omitempty"`                    // User's first name
	LastName  string    `json:"last_name,omitempty"`                     // User's last name
	Email     string    `json:"email,omitempty"`                         // User's email
	CreatedAt time.Time `json:"created_at"`                              // Creation timestamp
	UpdatedAt time.Time `json:"updated_at"`                              // Last update timestamp
}