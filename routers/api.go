package routers

import (
	"github.com/chirani/book-rest/services"
	"github.com/gin-gonic/gin"
)

func ApiRouterGroup(group *gin.RouterGroup) {

	group.GET("/api/books", services.FindBooks)
	group.POST("/api/books", services.CreateBook)
	group.GET("/api/books/:id", services.FindBook)
	group.PATCH("/api/books/:id", services.UpdateBook)
	group.DELETE("/api/books/:id", services.DeleteBook)

}
