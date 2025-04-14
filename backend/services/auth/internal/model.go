package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type AuthManager interface {
	LoginUser(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
}

type UserClient interface {
	GetUserByEmail(ctx context.Context, email string) (AuthDetials, error)
	UpdateUserLogin(ctx context.Context, email string) error
}

type Claims struct {
	Email    string `json:"email"`
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type AuthDetials struct {
	Email        string
	PasswordHash string
	Role         string
	TenantID     string
}
