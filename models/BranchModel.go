package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(255)"`
	Address     string   `gorm:"type:varchar(255)"`
	PhoneNumber string   `gorm:"type:varchar(255)"`
	Members     []Member `gorm:"foreignKey:BranchID"`
	Loans       []Loan   `gorm:"foreignKey:BranchID"`
	Rooms       []Room   `gorm:"foreignKey:BranchID"`
}
