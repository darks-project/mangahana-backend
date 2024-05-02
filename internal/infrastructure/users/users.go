package users

import (
	"api/internal/core"

	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type users struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *users {
	return &users{db: db}
}

func (r *users) AddSession(userId int) (string, error) {
	token, err := core.GenerateRandomString(128)
	if err != nil {
		return "", err
	}

	query := "INSERT INTO users.sessions (user_id, token) VALUES ($1, $2);"
	if err := r.db.QueryRowx(query, userId, token).Scan(); err != nil && err != pgx.ErrNoRows {
		return "", err
	}

	return token, nil
}

func (r *users) GetIdBySession(token string) (int, error) {
	var userId int
	query := "SELECT user_id FROM users.sessions WHERE token = $1"
	return userId, r.db.QueryRowx(query, token).Scan(&userId)
}

func (r *users) GetById(userId int) (core.User, error) {
	var output core.User
	query := "SELECT * FROM users.users WHERE id = $1"
	err := r.db.QueryRowx(query, userId).StructScan(&output)
	return output, err
}

func (r *users) GetByEmail(email string) (core.User, error) {
	var output core.User
	query := "SELECT * FROM users.users WHERE email = $1"
	err := r.db.QueryRowx(query, email).StructScan(&output)
	return output, err
}

func (r *users) HasPermission(userId int, permissionName string) error {
	user, err := r.GetById(userId)
	if err != nil {
		return err
	}

	query := "SELECT permissions FROM users.roles LEFT JOIN users.permissions ON permissions.id = any(roles.permissions) WHERE roles.id = $1;"

	return nil
}
