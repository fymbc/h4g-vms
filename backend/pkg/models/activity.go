package models

type Activity struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Datetime    string `gorm:"not null"`

	Registrations []ActivityRegistration `gorm:"many2many:activity_registrations;"`
}


