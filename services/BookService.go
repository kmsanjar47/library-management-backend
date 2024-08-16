package services

import (
	"main/config"
	"main/models"
	"main/repository"
)

func init() {
	config.ConnectToDB()
}

func FetchAllBooksService() []models.Book {

	books := repository.FetchAllBooksFromDB()

	return books
}

func CreateBook() models.Book {
	book := repository.AddNewBookToDB("Title", "ISBN", nil, 1, 1, 1)

	return book
}

func UpdateBook(bookID string, data map[string]interface{}) models.Book {
	book := repository.UpdateBookDataFromDB(bookID, data)

	return book
}

func DeleteBook(bookID string) {
	repository.DeleteBookFromDB(bookID)
}
