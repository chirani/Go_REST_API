package main

import (
	"net/http"

	"github.com/chirani/book-rest/models"
	"github.com/chirani/book-rest/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConnectDatabase()

	rgApi := r.Group("/api")

	routers.ApiRouterGroup(rgApi)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
