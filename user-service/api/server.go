package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", login)

	r.Use(authMiddleware())

	r.GET("/librarians", getAllLibrarian)
	r.GET("/librarians/:id", getLibrarian)
	r.POST("/librarians", adminMiddleware(), createLibrarian)
	r.PATCH("/librarians/:id", adminMiddleware(), updateLibrarian)
	r.DELETE("/librarians/:id", adminMiddleware(), deleteLibrarian)

	r.GET("/books", getAllBook)
	r.GET("/books/:id", getOneBook)
	r.GET("/books/description", searchBookDescription)
	r.GET("/books/:id/authors", getBookAuthors)
	r.POST("/books", librarianMiddleware(), createBook)
	r.DELETE("/books/:id", librarianMiddleware(), deleteBook)

	r.POST("/admins", adminMiddleware(), createAdmin)

	return r
}
