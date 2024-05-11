package services

import (
	"flea-market/dto"
	"flea-market/models"
	"flea-market/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(newUserInput dto.CreateUserInput) (*models.User, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func CreateNewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(newUserInput dto.CreateUserInput) (*models.User, error) {
	var newUser models.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser.Email = newUserInput.Email
	newUser.Password = string(hashedPassword)
	return s.repository.Signup(newUser)
}
