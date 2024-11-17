package domain

import (
	"context"
)

// RefreshTokenRequest represents the incoming request for refreshing tokens.
type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"` // Token provided by the client
}

// RefreshTokenResponse represents the response containing new access and refresh tokens.
type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`  // Newly generated access token
	RefreshToken string `json:"refreshToken"` // Newly generated refresh token
}

// RefreshTokenUsecase defines the interface for handling token refresh operations.
type RefreshTokenUsecase interface {
	// GetUserByID fetches the user from the database using the provided user ID.
	GetUserByID(c context.Context, id string) (User, error)

	// CreateAccessToken generates a new access token for the user.
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)

	// CreateRefreshToken generates a new refresh token for the user.
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)

	// ExtractIDFromToken validates the token and extracts the user ID from it.
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}