package repository

import "github.com/jmoiron/sqlx"

type Auth interface {
	Create()
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
