package service

import (
	"coffeecatalogue4"
	"coffeecatalogue4/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	salt       = "fjakjadlksa"
	signingKey = "sfjalfjlasjdlajkdf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user coffeecatalogue4.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(login string, password string) (string, error) {
	user, err := a.repo.GetUser(login, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		int(user.Id),
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%", hash.Sum([]byte(salt)))
}
