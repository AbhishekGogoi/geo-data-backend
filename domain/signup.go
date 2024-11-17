package domain

import (
	"context"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	// Create inserts a new user into the PostgreSQL database.
	Create(c context.Context, user *User) error

	// GetUserByEmail fetches a user by their email from the PostgreSQL database.
	GetUserByEmail(c context.Context, email string) (User, error)

	// CreateAccessToken generates a new access token for the user.
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)

	// CreateRefreshToken generates a new refresh token for the user.
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}