package repository

import (
	"debts/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) Auth {
	return &auth{db: db}
}

const userTable = "users"

func (a *auth) Create(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password) VALUES ($1, $2, $3) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *auth) GetByEmail() {

}

func (a *auth) UserByEmailExist() {

}
