package main

import (
	"log"
	"main/config"
	"main/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	config.ConnectToDB()
}

func main() {

	ADDRESS := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
	router := gin.Default()

	// Books
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)

	router.Run(ADDRESS)
}
