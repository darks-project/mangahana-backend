package application

import (
	"api/internal/application/users"
	"api/internal/infrastructure"
)

type UseCase struct {
	Users IUsers
}

func New(repo *infrastructure.Repo) *UseCase {
	return &UseCase{
		Users: users.New(repo.Users),
	}
}

type IUsers interface {
	LogIn(email, password string) (string, error)
	SignUp(email, password string) (string, error)

	GetOneById(userId int) map[string]interface{}
	GetIdByAccessToken(accessToken string) int
}
