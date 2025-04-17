package user

import (
	"context"
)

type UserManager interface {
	CreateUser(ctx context.Context, name, email, password string) (User, error)
	Login(ctx context.Context, email, password string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUserLogin(ctx context.Context, email string) error
}

type AuthClient interface {
	AuthenticateUser(ctx context.Context, email, password string) (string, error)
}

type Persistence interface {
	SaveUser(ctx context.Context, name, email, password string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUserLogin(ctx context.Context, email string) error
}

type User struct {
	Id           string
	Name         string
	Email        string
	PasswordHash string
	TenantID     string
	Role         string
}
