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

type Roastery interface {
	Create(userId int, roastery coffeecatalogue4.Roastery) (int, error)
}

type CoffeePack interface {
}

type Service struct {
	Authorization
	Roastery
	CoffeePack
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Roastery:      NewRoasteryService(repos.Roastery),
	}
}
