package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// createToken

func CreateToken(username string) (string, error) {
	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 3).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// validateToken
func VerifyExpirationTime(tokenTime time.Time) error {
	// claims, err := ParseToken(tokenString)
	// if err != nil {
	// 	return err
	// }

	expirationTime := tokenTime
	// time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return fmt.Errorf("token has expired")
	}

	return nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
// 			return
// 		}
// 		tokenString = tokenString[len("Bearer "):]

// 		if err := verifyToken(tokenString); err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			return
// 		}

// 		// Call the next handler
// 		c.Next()
// 	}
// }
