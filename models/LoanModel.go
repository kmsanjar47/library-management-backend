package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	BookID     uint   `gorm:"not null"`
	Book       Book   `gorm:"foreignKey:BookID"`
	MemberID   uint   `gorm:"not null"`
	Member     Member `gorm:"foreignKey:MemberID"`
	LoanDate   time.Time
	ReturnDate *time.Time
	Returned   bool
	Status     string `gorm:"type:varchar(255)"`
	BranchID   uint   `gorm:"not null"`
	Branch     Branch `gorm:"foreignKey:BranchID"`
}
