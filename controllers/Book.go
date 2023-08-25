package controllers

import (
	"github.com/chirani/book-rest/models"
)

func FindBooks() []models.Book {
	var books []models.Book
	models.DB.Find(&books)
	return books
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(input *CreateBookInput) *models.Book {
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	return &book
}

func FindBook(id string) *models.Book { // Get model if exist
	var book models.Book

	if err := models.DB.Where(id, id).First(&book).Error; err != nil {
		return nil
	}
	return &book
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(id string, input UpdateBookInput) *models.Book {

	var book models.Book
	if err := models.DB.Where(id, id).First(&book).Error; err != nil {
		return nil
	}

	models.DB.Model(&book).Updates(input)
	return &book

}

func DeleteBook(id string) {
	var book models.Book

	if err := models.DB.Where(id, id).First(&book).Error; err != nil {
		return
	}

	models.DB.Delete(&book)
}
