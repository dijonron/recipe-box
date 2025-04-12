package persistence

type User struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	CreadtedAt string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
	DeletedAt  string `db:"deleted_at"`
	LastLogin  string `db:"last_login"`
}
