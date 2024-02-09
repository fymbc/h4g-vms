package models

type ActivityRegistration struct {
	ID         int `gorm:"primary_key"`
	UserID     int `gorm:"not null"`
	ActivityID int `gorm:"not null"`
}
