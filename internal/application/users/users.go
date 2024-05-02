package users

import (
	"api/internal/infrastructure"

	"golang.org/x/crypto/bcrypt"
)

type users struct {
	repo infrastructure.IUsers
}

func New(repo infrastructure.IUsers) *users {
	return &users{repo: repo}
}

func (u *users) LogIn(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	hashedPassord, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), hashedPassord); err != nil {
		return "", err
	}

	return u.repo.AddSession(user.Id)
}
func (u *users) SignUp(email, password string) (string, error)

func (u *users) GetOneById(userId int) map[string]interface{}
func (u *users) GetIdByAccessToken(accessToken string) int
