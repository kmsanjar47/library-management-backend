package repository

import (
	"log"
	"main/config"
	"main/models"
	"time"
)

func init() {
	config.ConnectToDB()
}

func FetchAllBooksFromDB() []models.Book {

	books := []models.Book{}
	config.DB.Preload("Author").Preload("Publisher").Preload("Category").Find(&books)

	return books

}

func AddNewBookToDB(title string, isbn string, summery *string, authorID uint, publisherID uint, categoryID uint) models.Book {
	book := models.Book{
		Title:           title,
		Isbn:            isbn,
		Summary:         summery,
		AuthorID:        authorID,
		PublisherID:     publisherID,
		CategoryID:      categoryID,
		PublicationDate: time.Now(),
	}

	err := config.DB.Create(&book)
	if err != nil {
		log.Fatal(err)
	}
	return book

}

func UpdateBookDataFromDB(bookID string, data map[string]any) models.Book {
	books := models.Book{}
	err := config.DB.Find(&books, bookID).Error
	if err != nil {
		log.Fatal(err)

	}

	err_two := config.DB.Model(&books).Updates(data)
	if err_two != nil {
		log.Fatal(err_two)
	}

	return books

}

func DeleteBookFromDB(bookID string) {
	book := models.Book{}
	err := config.DB.Find(&book, bookID).Error
	if err != nil {
		log.Fatal(err)
	}

	err_two := config.DB.Delete(&book)
	if err_two != nil {
		log.Fatal(err_two)
	}

}
