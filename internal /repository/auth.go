package repository

import "github.com/jmoiron/sqlx"

type auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) Auth {
	return &auth{db: db}
}

func (a *auth) Create() {

}

func (a *auth) GetByEmail() {

}

func (a *auth) UserByEmailExist() {

}
