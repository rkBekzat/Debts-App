package service

import "debts/internal /repository"

type auth struct {
	repo repository.Auth
}

func NewAuth(repo repository.Auth) Auth {
	return &auth{repo: repo}
}
