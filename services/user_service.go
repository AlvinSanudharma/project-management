package services

import (
	"errors"

	"github.com/AlvinSanudharma/project-management/models"
	"github.com/AlvinSanudharma/project-management/repositories"
	"github.com/AlvinSanudharma/project-management/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("Email already registered")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}
