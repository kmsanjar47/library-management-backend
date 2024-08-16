package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Status      string `gorm:"type:varchar(255)"`
	BookedBy    uint   `gorm:"not null"`
	Member      Member `gorm:"foreignKey:BookedBy"`
	BranchID    uint   `gorm:"not null"`
	Branch      Branch `gorm:"foreignKey:BranchID"`
	CurrentSlot time.Time
}
