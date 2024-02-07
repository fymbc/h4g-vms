package users

import (
	"h4g-vms/pkg/models"

	"gorm.io/gorm"
)



func Read(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func Create(db *gorm.DB, user *models.User) (*models.User, error) {
	result := db.Create(user).First(user)
	return user, result.Error
}
