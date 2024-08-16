package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string    `gorm:"type:varchar(255)"`
	Isbn            string    `gorm:"type:varchar(255)"`
	Summary         *string   `gorm:"type:text"`
	AuthorID        uint      `gorm:"not null"`
	Author          Author    `gorm:"foreignKey:AuthorID"`
	PublisherID     uint      `gorm:"not null"`
	Publisher       Publisher `gorm:"foreignKey:PublisherID"`
	CategoryID      uint      `gorm:"not null"`
	Category        Category  `gorm:"foreignKey:CategoryID"`
	PublicationDate time.Time
}
