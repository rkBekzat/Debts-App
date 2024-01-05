package service

import "debts/internal /repository"

type Auth interface {
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Auth: NewAuth(repo.Auth),
	}
}
