package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(255)"`
	LastName    string `gorm:"type:varchar(255)"`
	DateOfBirth time.Time
	DateOfDeath *time.Time
	Books       []Book `gorm:"foreignKey:AuthorID"`
}
