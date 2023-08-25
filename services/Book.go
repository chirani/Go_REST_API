package services

import (
	"net/http"

	"github.com/chirani/book-rest/controllers"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var input controllers.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := controllers.CreateBook(&input)
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func FindBooks(c *gin.Context) {
	books := controllers.FindBooks()

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	id := c.Param("id")
	book := controllers.FindBook(id)

	if book != nil {
		c.JSON(http.StatusOK, gin.H{"data": book})
	}
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var input controllers.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := controllers.UpdateBook(id, input)
	if book != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	controllers.DeleteBook(id)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
