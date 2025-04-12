package user

import (
	"context"
)

type UserManager interface {
	CreateUser(ctx context.Context, name, email, password string) (string, error)
	GetUserDetails(ctx context.Context, userID string) (User, error)
}

type AuthClient interface {
	AuthenticateUser(ctx context.Context, email, password string) (bool, error)
}

type Persistence interface {
	SaveUser(ctx context.Context) error
	GetUserByID(ctx context.Context, userID string) (User, error)
	UpdateUser(ctx context.Context, userID string) error
	DeleteUser(ctx context.Context, userID string) error
}

type User struct {
	ID    string
	Name  string
	Email string
}
