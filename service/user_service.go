package service

import (
	"sea-study/api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}

func GetUserByEmail(db *gorm.DB, user *models.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func GetUserByID(db *gorm.DB, user *models.User, userID uuid.UUID) error {
	return db.Where("user_id = ?", userID).First(user).Error
}