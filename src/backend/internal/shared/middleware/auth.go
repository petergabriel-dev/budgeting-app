package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petergabriel-dev/budgeting-app/internal/features/auth"
)

// SessionAuth creates a middleware that validates session tokens
func SessionAuth(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(auth.SessionCookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			return
		}

		user, err := authService.ValidateSession(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
			return
		}

		// Store user in context for handlers
		c.Set("user", user)
		c.Set("userID", user.ID)

		c.Next()
	}
}

// OptionalSessionAuth creates a middleware that validates session tokens but doesn't require them
// Useful for endpoints that can work with or without authentication
func OptionalSessionAuth(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(auth.SessionCookieName)
		if err != nil {
			c.Next()
			return
		}

		user, err := authService.ValidateSession(c.Request.Context(), token)
		if err != nil {
			c.Next()
			return
		}

		// Store user in context for handlers
		c.Set("user", user)
		c.Set("userID", user.ID)

		c.Next()
	}
}
