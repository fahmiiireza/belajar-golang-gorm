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

	r.POST("/admin", createAdmin)

	r.POST("/librarian", createLibrarian)
	r.GET("/librarian/:id", getLibrarian)
	r.PATCH("/librarian/:id", updateLibrarian)
	r.DELETE("/librarian/:id", deleteLibrarian)
	r.GET("/search", searchHandler)

	r.GET("/book", getAllBook)
	r.GET("/book/:id", getOneBook)
	return r
}
