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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set the username in the context
		c.Set("username", username)

		c.Next()
	}
}
