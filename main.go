package main

import (
	"fmt"

	"github.com/Man4ct/belajar-golang-gorm/controllers/concur"
	controllers "github.com/Man4ct/belajar-golang-gorm/controllers/users"
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/user", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.POST("/librarian", controllers.CreateLibrarian)
	r.GET("/librarian/:id", controllers.GetLibrarian)
	r.PATCH("/librarian/:id", controllers.UpdateLibrarian)

	r.GET("/search", concur.SearchHandler)

	// Start the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
