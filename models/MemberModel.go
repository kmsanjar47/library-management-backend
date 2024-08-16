package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	FirstName      string  `gorm:"type:varchar(255)"`
	LastName       string  `gorm:"type:varchar(255)"`
	Email          string  `gorm:"type:varchar(255)"`
	PhoneNumber    string  `gorm:"type:varchar(255)"`
	Address        *string `gorm:"type:varchar(255)"`
	MembershipDate time.Time
	BranchID       uint   `gorm:"not null"`
	Branch         Branch `gorm:"foreignKey:BranchID"`
	Loans          []Loan `gorm:"foreignKey:MemberID"`
	Rooms          []Room `gorm:"foreignKey:BookedBy"`
}
