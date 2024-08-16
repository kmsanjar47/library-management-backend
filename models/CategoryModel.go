package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Books []Book `gorm:"foreignKey:CategoryID"`
}
