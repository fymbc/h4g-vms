package models

type User struct {
	ID                int    `gorm:"not null"`
	Email             string `gorm:"not null;unique"`
	EncryptedPassword string `gorm:"not null"`

	RegisteredActivities []Activity `gorm:"many2many:users_activities;"`
}

