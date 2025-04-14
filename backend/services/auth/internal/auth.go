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
	INVALID_CREDENTIALS    string = "invalid credentials"
	INVALID_SIGNING_METHOD string = "unexpected signing method"
	FAILED_TO_GEN_JWT      string = "failed to generate jwt"
	INVALID_JWT            string = "dailed to validate jwt"
)

var (
	errInvalidCredentials   = errors.New(INVALID_CREDENTIALS)
	errInvalidJWT           = errors.New(INVALID_JWT)
	errInvalidSigningMethod = errors.New(INVALID_SIGNING_METHOD)
	errFailedToGenJwt       = errors.New(FAILED_TO_GEN_JWT)
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
	authDetials, err := a.userClient.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Verify password
	slog.Debug("p", "password", authDetials.PasswordHash)
	if err := bcrypt.CompareHashAndPassword([]byte(authDetials.PasswordHash), []byte(password)); err != nil {
		slog.Info(INVALID_CREDENTIALS, "error", err)
		return "", errInvalidCredentials
	}

	// Generate JWT token
	token, err := a.generateJWT(email, authDetials.TenantID, authDetials.Role)
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

func (a *authManager) ValidateToken(ctx context.Context, token string) (bool, error) {
	valid, err := a.validateToken(token)
	if err != nil {
		return false, err
	}

	return valid, nil
}

// generateJWT generates a JWT token for the authenticated user
func (a *authManager) generateJWT(email, tenantID, role string) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(time.Duration(a.jwtExpirationHours) * time.Hour)

	// Create claims
	claims := &Claims{
		email:    email,
		tenantID: tenantID,
		role:     role,
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

// validateToken validates the JWT token and returns the claims if valid
func (a *authManager) validateToken(token string) (bool, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Warn(INVALID_SIGNING_METHOD, "method", token.Header["alg"])
			return nil, errInvalidSigningMethod
		}
		return []byte(a.jwtSecret), nil
	})
	if err != nil {
		slog.Info(INVALID_JWT, "error", err)
		return false, errInvalidJWT
	}

	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		slog.Debug("token validated successfully", "email", claims.email, "tenantID", claims.tenantID, "role", claims.role)
		return true, nil
	}

	slog.Info(INVALID_JWT)
	return false, errInvalidJWT
}
