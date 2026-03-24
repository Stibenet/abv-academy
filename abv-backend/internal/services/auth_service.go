package services

import (
	"abv-backend/internal/models"
	"abv-backend/internal/repository"
	"abv-backend/internal/utils"
	"errors"
)

type AuthService struct {
	repo       *repository.UserRepository
	jwtSecret  string
	jwtTTLHour int
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string, jwtTTLHour int) *AuthService {
	return &AuthService{
		repo:       repo,
		jwtSecret:  jwtSecret,
		jwtTTLHour: jwtTTLHour,
	}
}

func (s *AuthService) Register(user *models.User, rawPassword string) error {
	hash, err := utils.HashPassword(rawPassword)
	if err != nil {
		return err
	}
	user.Password = hash
	return s.repo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, string(user.Role), s.jwtSecret, s.jwtTTLHour)
	if err != nil {
		return "", err
	}

	return token, nil
}
