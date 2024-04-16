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

	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		user, err := helper.CreateUser(tx, newAdmin.Username, newAdmin.Email, newAdmin.Password, newAdmin.FullName, model.RoleAdmin)
		if err != nil {
			return err
		}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin created successfully"})
}
