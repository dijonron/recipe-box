package persistence

import (
	user "github.com/dijonron/recipe-box/services/user/internal"
)

type User struct {
	Name       string `db:"name"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	CreadtedAt string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	DeletedAt  string `db:"deleted_at"`
	LastLogin  string `db:"last_login"`
}

func toUser(u User) user.User {
	return user.User{
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.Password,
	}
}
