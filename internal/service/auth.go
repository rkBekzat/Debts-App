package service

import (
	"crypto/sha1"
	"debts/internal/model"
	"debts/internal/repository"
	"fmt"
)

type auth struct {
	repo repository.Auth
}

const salty = "eqwrghyttr132tgrbr132rg"

func NewAuth(repo repository.Auth) Auth {
	return &auth{repo: repo}
}

func (a *auth) SignUp(user model.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return a.repo.Create(user)
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salty)))
}
