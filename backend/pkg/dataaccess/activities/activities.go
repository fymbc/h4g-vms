package activities

import (
	"h4g-vms/pkg/models"

	"gorm.io/gorm"
)

func List(db *gorm.DB) ([]models.Activity, error) {
	var activities []models.Activity
	result := db.Find(&activities)
	return activities, result.Error
}

func Read(db *gorm.DB, activityId int) (*models.Activity, error) {
	var activity models.Activity
	result := db.Where("id = ?", activityId).First(&activity)
	return &activity, result.Error
}

func Create(db *gorm.DB, activity *models.Activity) (*models.Activity, error) {
	result := db.Create(activity).First(activity)
	return activity, result.Error
}

func Update(db *gorm.DB, activity *models.Activity) (*models.Activity, error) {
	result := db.Save(activity).First(activity)
	return activity, result.Error
}

func Delete(db *gorm.DB, activityId int) error {
	result := db.Delete(&models.Activity{}, activityId)
	return result.Error
}

