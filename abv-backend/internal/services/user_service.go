package services

import (
	"abv-backend/internal/models"
	"abv-backend/internal/repository"
	"abv-backend/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *models.User, rawPassword string) error {
	hash, err := utils.HashPassword(rawPassword)
	if err != nil {
		return err
	}
	user.Password = hash
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(id uint, input *models.User, rawPassword *string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Role = input.Role

	if rawPassword != nil && *rawPassword != "" {
		hash, err := utils.HashPassword(*rawPassword)
		if err != nil {
			return nil, err
		}
		user.Password = hash
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
