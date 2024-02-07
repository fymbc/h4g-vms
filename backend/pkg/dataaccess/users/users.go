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
