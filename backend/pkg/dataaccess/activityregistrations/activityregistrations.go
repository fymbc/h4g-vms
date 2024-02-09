package activityregistrations

import (
	"h4g-vms/pkg/models"

	"gorm.io/gorm"
)

func Create(db *gorm.DB, activityRegistration *models.ActivityRegistration) error {
	result := db.Create(activityRegistration)
	return result.Error
}

func Delete(db *gorm.DB, activityRegistrationID int) error {
	result := db.Delete(&models.ActivityRegistration{}, activityRegistrationID)
	return result.Error
}
