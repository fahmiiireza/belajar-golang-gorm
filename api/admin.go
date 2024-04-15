package api

import (
	"net/http"

	"github.com/Man4ct/belajar-golang-gorm/db"
	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"github.com/Man4ct/belajar-golang-gorm/helper"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func createAdmin(c *gin.Context) {
	var newAdmin AdminRequest

	if err := c.BindJSON(&newAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use automatic transaction with callback function
	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		// Create user
		user, err := helper.CreateUser(tx, newAdmin.Username, newAdmin.Email, newAdmin.Password, newAdmin.FullName)
		if err != nil {
			return err
		}

		// Create librarian
		admin := model.Admin{
			Salary:           newAdmin.Salary,
			EmploymentStatus: newAdmin.EmploymentStatus,
			UserID:           user.ID,
		}
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create librarian"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Librarian created successfully"})
}
