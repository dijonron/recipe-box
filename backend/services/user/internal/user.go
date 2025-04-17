package user

import (
	"context"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

var (
	INVALID_CREDENTIALS string = "invalid credentials"
)

var (
	errInvalidCredentials = errors.New(INVALID_CREDENTIALS)
)

type userManager struct {
	persistence Persistence
	authclient  AuthClient
}

var _ UserManager = (*userManager)(nil)

func NewUserManager(p Persistence, ac AuthClient) UserManager {
	return &userManager{
		persistence: p,
		authclient:  ac,
	}
}

func (u *userManager) CreateUser(ctx context.Context, name, email, password string) (User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return User{}, err
	}

	userID, err := u.persistence.SaveUser(ctx, name, email, hashedPassword)
	if err != nil {
		return User{}, err
	}

	u.persistence.UpdateUserLogin(ctx, email)

	user := User{
		Id:    userID,
		Name:  name,
		Email: email,
		Role:  "user", // hardcode for now, since new users will always have the user role
	}

	return user, nil
}

func (u *userManager) Login(ctx context.Context, email, password string) (User, error) {
	user, err := u.persistence.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		slog.Info(INVALID_CREDENTIALS, "error", err)
		return User{}, errInvalidCredentials
	}

	u.persistence.UpdateUserLogin(ctx, email)

	return user, nil
}

func (u *userManager) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := u.persistence.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *userManager) UpdateUserLogin(ctx context.Context, email string) error {
	err := u.persistence.UpdateUserLogin(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

// hashPassword hashes the given plaintext password using bcrypt with the default cost.
// It returns the hashed password as a string or an error if the hashing process fails.
//
// Parameters:
//   - password: The plaintext password to be hashed.
//
// Returns:
//   - A string containing the hashed password.
//   - An error if the password hashing fails.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		return "", err
	}

	return string(hashedPassword), nil
}
