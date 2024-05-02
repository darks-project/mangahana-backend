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
