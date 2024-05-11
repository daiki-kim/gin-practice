package repositories

import (
	"flea-market/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(newUser models.User) error // repositoriesではSigninよりCreateUserの方がわかりやすい
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
