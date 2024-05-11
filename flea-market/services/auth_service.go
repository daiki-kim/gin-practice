package services

import (
	"flea-market/models"
	"flea-market/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error // dtoの構造体よりフィールドを入れる方が読みやすい
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) SignUp(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err // errにstring型のエラー内容がそのまま入っているためerr.Error()ではなくerrのまま返す
	}
	newUser := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(newUser)
}
