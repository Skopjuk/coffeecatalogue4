package service

import (
	"coffeecatalogue4"
	"coffeecatalogue4/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "fjakjadlksa"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user coffeecatalogue4.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%", hash.Sum([]byte(salt)))
}
