package persistence

import (
	"database/sql"

	user "github.com/dijonron/recipe-box/services/user/internal"
)

type User struct {
	Id           string         `db:"id"`
	Name         string         `db:"name"`
	Email        string         `db:"email"`
	PasswordHash string         `db:"password_hash"`
	TenantID     sql.NullString `db:"tenant_id"`
	Role         string         `db:"role"`
	LastLogin    string         `db:"last_login"`
	CreadtedAt   string         `db:"created_at"`
	UpdatedAt    string         `db:"updated_at"`
	DeletedAt    string         `db:"deleted_at"`
}

func toUser(u User) user.User {
	return user.User{
		Id:           u.Id,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
		TenantID:     u.TenantID.String,
	}
}
