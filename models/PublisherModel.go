package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name    string  `gorm:"type:varchar(255)"`
	Address *string `gorm:"type:varchar(255)"`
	Website *string `gorm:"type:varchar(255)"`
	Books   []Book  `gorm:"foreignKey:PublisherID"`
}
