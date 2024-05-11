package repositories

import (
	"flea-market/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	Signup(newUser models.User) (*models.User, error)
}

type AuthRepository struct {
	users *gorm.DB
}

func CreateNewAuthRepository(users *gorm.DB) IAuthRepository {
	return &AuthRepository{users: users}
}

func (r *AuthRepository) Signup(newUser models.User) (*models.User, error) {
	result := r.users.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}
