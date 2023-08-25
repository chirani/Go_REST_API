/*
source : https://blog.logrocket.com/rest-api-golang-gin-gorm/
*/

package main

import (
	"net/http"

	"github.com/chirani/book-rest/controllers"
	"github.com/chirani/book-rest/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/books", controllers.FindBooks)
	r.POST("/api/books", controllers.CreateBook)
	r.GET("/api/books/:id", controllers.FindBook)
	r.PATCH("/api/books/:id", controllers.UpdateBook)
	r.DELETE("/api/books/:id", controllers.DeleteBook)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
