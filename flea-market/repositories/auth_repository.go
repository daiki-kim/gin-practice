package repositories

import (
	"flea-market/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(newUser models.User) error // repositoriesではSigninよりCreateUserの方がわかりやすい
	FindUserByEmail(email string) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(newUser models.User) error {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var foundUser models.User
	result := r.db.First(&foundUser, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundUser, nil
}
