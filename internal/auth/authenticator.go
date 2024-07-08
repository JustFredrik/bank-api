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
}

func Authenticator(required_role string) (c gin.HandlerFunc) {

	return func(c *gin.Context) {

		// Validate that the API key has Access to resource
		if !KeyHasAccess(c, required_role) {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		c.Next()
	}
}
