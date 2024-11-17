package domain

import (
	"context"
)

// LoginRequest represents the incoming request for user login.
type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"` // Email provided by the user
	Password string `form:"password" binding:"required"`   // Password provided by the user
}

// LoginResponse represents the response containing the access and refresh tokens.
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`  // Newly generated access token
	RefreshToken string `json:"refreshToken"` // Newly generated refresh token
}

// LoginUsecase defines the interface for user login operations.
type LoginUsecase interface {
	// GetUserByEmail fetches the user by email from the database.
	GetUserByEmail(c context.Context, email string) (User, error)

	// CreateAccessToken generates a new access token for the user.
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)

	// CreateRefreshToken generates a new refresh token for the user.
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}