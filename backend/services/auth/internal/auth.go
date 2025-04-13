package auth

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/dijonron/recipe-box/services/auth/cmd/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	INVALID_CREDENTIALS string = "invalid credentials"
	FAILED_TO_GEN_JWT   string = "failed to generate jwt"
)

var (
	errInvalidCredentials = errors.New(INVALID_CREDENTIALS)
	errFailedToGenJwt     = errors.New(FAILED_TO_GEN_JWT)
)

type authManager struct {
	jwtSecret          string
	jwtExpirationHours int
	userClient         UserClient
}

var _ AuthManager = (*authManager)(nil)

func NewAuthManager(cfg config.AuthConfig, uc UserClient) AuthManager {
	return &authManager{
		jwtSecret:          cfg.JWTSecret,
		jwtExpirationHours: cfg.JWTExpiration,
		userClient:         uc,
	}
}

func (a *authManager) LoginUser(ctx context.Context, email, password string) (string, error) {
	// Get user
	user, err := a.userClient.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Verify password
	slog.Debug("p", "password", user.PasswordHash)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		slog.Info(INVALID_CREDENTIALS, "error", err)
		return "", errInvalidCredentials
	}

	// Generate JWT token
	token, err := a.generateJWT(email)
	if err != nil {
		slog.Error(FAILED_TO_GEN_JWT, "error", err)
		return "", errFailedToGenJwt
	}

	// Update last login timestamp
	if err := a.userClient.UpdateUserLogin(ctx, email); err != nil {
		slog.Error("failed to update user login", "error", err)
	}

	return token, nil
}

// generateJWT generates a JWT token for the authenticated user
func (a *authManager) generateJWT(email string) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(time.Duration(a.jwtExpirationHours) * time.Hour)

	// Create claims
	claims := &Claims{
		email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "recipe-box-auth-service",
			Subject:   email,
			ID:        "",
			Audience:  []string{"recipe-box-services"},
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
