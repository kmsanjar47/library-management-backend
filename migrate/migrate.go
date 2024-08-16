package main

import (
	"main/config"
	"main/models"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Author{}, &models.Category{}, &models.Publisher{}, &models.Book{}, &models.Branch{}, &models.Member{}, &models.Loan{}, &models.Room{})

}
