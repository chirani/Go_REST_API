package main

import (
	"net/http"

	"github.com/chirani/book-rest/models"
	"github.com/chirani/book-rest/services"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/books", services.FindBooks)
	r.POST("/api/books", services.CreateBook)
	r.GET("/api/books/:id", services.FindBook)
	r.PATCH("/api/books/:id", services.UpdateBook)
	r.DELETE("/api/books/:id", services.DeleteBook)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
