// api/server.go

package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", login)

	// Use the middleware for protected routes
	r.Use(authMiddleware())

	r.GET("/librarians", getAllLibrarian)
	r.GET("/librarian/:id", getLibrarian)
	r.POST("/librarian", adminMiddleware(), createLibrarian)
	r.PATCH("/librarian/:id", adminMiddleware(), updateLibrarian)
	r.DELETE("/librarian/:id", adminMiddleware(), deleteLibrarian)

	r.GET("/books", getAllBook)
	r.GET("/book/:id", getOneBook)
	r.GET("/book/description", searchBookDescription)
	r.GET("/book/:id/authors", getBookAuthors)

	r.POST("/admin", adminMiddleware(), createAdmin)

	return r
}
