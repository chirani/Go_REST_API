package controllers

import (
	"github.com/chirani/book-rest/models"
	"gorm.io/gorm"
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

func CreateBook(input *CreateBookInput) *gorm.DB {
	book := models.Book{Title: input.Title, Author: input.Author}

	return models.DB.Create(&book)
}

func FindBook(id string) *gorm.DB {
	var book models.Book

	return models.DB.Where(id, id).First(&book)
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(id string, input UpdateBookInput) *gorm.DB {

	var book models.Book
	if err := models.DB.Where(id, id).First(&book).Error; err != nil {
		return nil
	}

	return models.DB.Model(&book).Updates(input)

}

func DeleteBook(id string) {
	var book models.Book

	if err := models.DB.Where(id, id).First(&book).Error; err != nil {
		return
	}

	models.DB.Delete(&book)
}
