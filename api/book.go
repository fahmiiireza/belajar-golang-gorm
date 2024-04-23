package api

import (
	"net/http"

	"github.com/Man4ct/belajar-golang-gorm/db"
	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"github.com/Man4ct/belajar-golang-gorm/helper"
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

	if result.Error != nil {
		if helper.IsNotFound(result.Error) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Books not found"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func getOneBook(c *gin.Context) {
	var book model.Book
	result := db.GetDB().First(&book, c.Param("id"))
	// .Preload("Authors")
	if result.Error != nil {
		if helper.IsNotFound(result.Error) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}
	}

	db.GetDB().Model(&book).Association("Authors").Find(&book.Authors)

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func getBookAuthors(c *gin.Context) {
	var book model.Book

	if err := db.GetDB().Preload("Authors").First(&book, c.Param("id")).Error; err != nil {
		if helper.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Books not found"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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
	if err := db.GetDB().Model(&model.Book{}).Where("isbn = ?", bookRequest.ISBN).First(&model.Book{}).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ISBN already exists"})
		return
	}

	var librarianID uint
	if err := db.GetDB().Model(&model.Librarian{}).Where("user_id = ?", userIDCreatedBy).Select("id").First(&librarianID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find librarian ID"})
		return
	}

	var book = model.Book{
		Title:       bookRequest.Title,
		ISBN:        bookRequest.ISBN,
		Description: bookRequest.Description,
		Language:    bookRequest.Language,
		TotalCopy:   bookRequest.TotalCopy,
		ShelfID:     bookRequest.ShelfID,
		CategoryID:  bookRequest.CategoryID,
		CreatedBy:   librarianID,
	}

	// Check if ShelfID is provided
	if bookRequest.ShelfID != 0 {
		book.ShelfID = uint(bookRequest.ShelfID)
	}

	// Check if CategoryID is provided
	if bookRequest.CategoryID != 0 {
		book.CategoryID = uint(bookRequest.CategoryID)
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

	c.JSON(http.StatusNoContent, nil)
}
