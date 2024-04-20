package api

import (
	"net/http"

	"github.com/Man4ct/belajar-golang-gorm/db"
	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"github.com/gin-gonic/gin"
)

func getAllBook(c *gin.Context) {
	var books []model.Book
	result := db.GetDB().Preload("Authors").Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func getOneBook(c *gin.Context) {
	var book model.Book
	result := db.GetDB().First(&book, c.Param("id"))
	// .Preload("Authors")
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	db.GetDB().Model(&book).Association("Authors").Find(&book.Authors)

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func getBookAuthors(c *gin.Context) {
	var book model.Book

	if err := db.GetDB().Preload("Authors").First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find the book",
		})
		return
	}

	var authors []model.Author
	authors = append(authors, book.Authors...)
	// for _, author := range book.Authors {
	// 	authors = append(authors, author)
	// }

	c.JSON(http.StatusOK, gin.H{"authors": authors})
}

func createBook(c *gin.Context) {
	var bookRequest BookRequest

	if err := c.BindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usnLoggedIn, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User logged in info not found in context"})
		return
	}
	var userIDCreatedBy uint
	if err := db.GetDB().Model(&model.User{}).Where("username = ? AND users.deleted_at IS NULL", usnLoggedIn).Select("id").First(&userIDCreatedBy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user logged in"})
		return
	}

	var librarianID uint
	if err := db.GetDB().Model(&model.Admin{}).Where("user_id = ?", userIDCreatedBy).Select("id").First(&librarianID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find librarian ID"})
		return
	}

	var book = model.Book{
		Title:       bookRequest.Title,
		ISBN:        bookRequest.ISBN,
		Description: bookRequest.Description,
		Language:    bookRequest.Language,
		TotalCopy:   bookRequest.TotalCopy,
		ShelfID:     uint(bookRequest.ShelfID),
		CategoryID:  uint(bookRequest.CategoryID),
		CreatedBy:   librarianID,
	}

	if err := db.GetDB().Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": book})

}

func deleteBook(c *gin.Context) {
	var book model.Book
	if err := db.GetDB().First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find the book"})
		return
	}

	if err := db.GetDB().Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the book"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Book deleted"})
}
