// Package auth provides authentication utilities and moddleware for the server.
package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// requestContainsToken checks if a
func requestContainsToken(c *gin.Context) bool {
	return true
}

func CreateMockKeys() {
	NewAPIKey(ROLE_ADMIN, 0)
	NewAPIKey(ROLE_ACCOUNT, 54400001111)
	NewAPIKey(ROLE_ACCOUNT, 13371337984)
}

func Authenticator(required_role string) (c gin.HandlerFunc) {

	return func(c *gin.Context) {

		// Validate that the API key format
		hasAccess, err := KeyHasAccess(c, required_role)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
			c.Abort()
			return
		}
		// Check if Key has Access
		if !hasAccess {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Your API key is not authorized to access the requested resource."})
			c.Abort()
			return
		}

		// User Authentication has passed, move on to next middleware / handler.
		c.Next()
	}
}
