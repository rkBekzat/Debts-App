package service

import (
	"debts/internal/model"
	"debts/internal/repository"
)

type Auth interface {
	SignUp(user model.User) (int, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Auth: NewAuth(repo.Auth),
	}
}
