package api

import (
	"net/http"
	"time"

	"github.com/Man4ct/belajar-golang-gorm/helper"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims, err := helper.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if err := helper.VerifyExpirationTime(expirationTime); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		}

		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Role not found in context"})
			return
		}
		if role != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}

func librarianMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Role not found in context"})
			return
		}
		if role != "ADMIN" && role != "LIBRARIAN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
