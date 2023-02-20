package repository

import (
	"coffeecatalogue4"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user coffeecatalogue4.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (coffeecatalogue4.User, error) {
	var user coffeecatalogue4.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
