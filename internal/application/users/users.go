package users

import (
	"api/internal/core"
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

func (u *users) SignUp(email, password string) (string, error) {
	return "", nil
}

func (u *users) GetOneById(userId int) (map[string]interface{}, error) {
	output := make(map[string]interface{})

	user, err := u.repo.GetById(userId)
	if err != nil {
		return nil, err
	}

	output["id"] = user.Id
	output["username"] = user.Username
	output["photo"] = user.Photo
	output["description"] = user.Description

	return output, nil
}

func (u *users) GetBySession(token string) (core.User, error) {
	var user core.User

	id, err := u.repo.GetIdBySession(token)
	if err != nil {
		return user, err
	}

	return u.repo.GetById(id)
}
