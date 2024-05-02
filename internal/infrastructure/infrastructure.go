package infrastructure

import (
	"api/internal/core"
	"api/internal/infrastructure/users"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	Users IUsers
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		Users: users.New(db),
	}
}

type IUsers interface {
	AddSession(userId int) (string, error)

	HasPermission(userId int, permissionName string) error

	GetIdBySession(token string) (int, error)
	GetById(userId int) (core.User, error)
	GetByEmail(email string) (core.User, error)
}
