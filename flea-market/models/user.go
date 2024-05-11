package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" binding:"email"`
	Password string `gorm:"not null;min=8"`
}
