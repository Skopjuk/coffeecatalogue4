package service

import (
	"coffeecatalogue4"
	"coffeecatalogue4/pkg/repository"
)

type Authorization interface {
	CreateUser(user coffeecatalogue4.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
