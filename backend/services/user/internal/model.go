package user

import (
	"context"
)

type UserManager interface {
	CreateUser(ctx context.Context, name, email, password string) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUserLogin(ctx context.Context, email string) error
}

type AuthClient interface {
	AuthenticateUser(ctx context.Context, email, password string) (string, error)
}

type Persistence interface {
	SaveUser(ctx context.Context, name, email, password string) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUserLogin(ctx context.Context, email string) error
	DeleteUser(ctx context.Context, userID string) error
}

type User struct {
	Name         string
	Email        string
	PasswordHash string
}
