package user

import "context"

type userManager struct {
	persistence Persistence
}

var _ UserManager = (*userManager)(nil)

func NewUserManager(p Persistence) UserManager {
	return &userManager{
		persistence: p,
	}
}

func (u *userManager) CreateUser(ctx context.Context, name, email, password string) (string, error) {
	// Implementation for creating a user
	return "", nil
}

func (u *userManager) GetUserDetails(ctx context.Context, userID string) (User, error) {
	// Implementation for getting user details
	return User{}, nil
}
