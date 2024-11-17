package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtCustomClaims defines the custom claims for the access token.
type JwtCustomClaims struct {
	ID   string `json:"id"`   // UUID of the user
	Name string `json:"name"` // Name of the user
	jwt.RegisteredClaims
}

// JwtCustomRefreshClaims defines the custom claims for the refresh token.
type JwtCustomRefreshClaims struct {
	ID string `json:"id"` // UUID of the user
	jwt.RegisteredClaims
}