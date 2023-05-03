package service

import (
	"crypto/sha1"
	"fmt"
	"todo"
	"todo/pkg/repository"
)

const salt = "arynkrasavchik"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePassword(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
