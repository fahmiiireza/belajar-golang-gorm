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
	var adminRequest AdminRequest
	var admin model.Admin

	if err := c.BindJSON(&adminRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if adminRequest.EmploymentStatus == "RESIGNED" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employment status cannot be resigned"})
		return
	}

	if !helper.IsValidEmail(adminRequest.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if exists, err := helper.CheckExistingUser(adminRequest.Username, adminRequest.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check for existing user"})
		return
	} else if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that username or email already exists"})
		return
	}

	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := helper.CreateUser(tx, adminRequest.Username, adminRequest.Email, adminRequest.Password, adminRequest.FullName, model.RoleAdmin)
		if err != nil {
			return err
		}

		admin = model.Admin{
			Salary:           adminRequest.Salary,
			EmploymentStatus: adminRequest.EmploymentStatus,
			UserID:           user.ID,
		}
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"admin": admin})
}
