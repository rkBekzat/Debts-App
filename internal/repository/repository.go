package repository

import (
	"debts/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	Create(user model.User) (int, error)
	GetByEmail()
	UserByEmailExist()
}

type Repo struct {
	Auth
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Auth: NewAuth(db),
	}
}
