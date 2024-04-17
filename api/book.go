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
