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
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func getOneBook(c *gin.Context) {
	var book model.Book
	result := db.GetDB().Preload("Authors").First(&book, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

// var librarian model.Librarian
// result := db.GetDB().Preload("User").Where("employment_status != ?", "RESIGNED").First(&librarian, c.Param("id"))
// if result.Error != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message": result.Error,
// 	})
// 	return
// }
// c.JSON(http.StatusOK, gin.H{
// 	"librarian": librarian,
// })
