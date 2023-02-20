package repository

import (
	"coffeecatalogue4"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user coffeecatalogue4.User) (int, error)
	GetUser(login, password string) (coffeecatalogue4.User, error)
}

type Coffee interface {
}

type Roastery interface {
}

type Repository struct {
	Authorization
	Coffee
	Roastery
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
