package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Parse and validate JWT token
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		// Use a secret key for verification
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Extract user information from the token
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(uint32) // Assuming you store user ID in the token

	// Make user information available to subsequent handlers
	c.Set("userID", userID)
}
