package controllers

import (
	"main/services"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := services.FetchAllBooksService()

	response := utils.NewResponse("success", "Books data successfully retrieved", books)

	c.JSON(http.StatusOK, response)
}

func CreateBook(c *gin.Context) {
	book := services.CreateBook()

	response := utils.NewResponse("success", "Book successfully created", book)

	c.JSON(http.StatusOK, response)
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("id")
	data := map[string]interface{}{
		"title": "New Title",
	}

	book := services.UpdateBook(bookID, data)

	response := utils.NewResponse("success", "Book successfully updated", book)

	c.JSON(http.StatusOK, response)
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	services.DeleteBook(bookID)

	response := utils.NewResponse("success", "Book successfully deleted", nil)

	c.JSON(http.StatusOK, response)
}
