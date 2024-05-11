package services

import (
	"errors"
	"flea-market/models"
	"flea-market/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error // dtoの構造体よりフィールドを入れる方が読みやすい
	LogIn(email string, password string) (*string, error)
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

func (s *AuthService) LogIn(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUserByEmail(email)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		return nil, err
	}
	return CreateToken(foundUser.ID, foundUser.Email)
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
