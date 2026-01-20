package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petergabriel-dev/budgeting-app/internal/features/auth"
)

// CSRF creates a middleware that validates CSRF tokens using the Double Submit Cookie pattern
// The client must send the CSRF token in the X-CSRF-Token header, matching the csrf_token cookie
func CSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip CSRF validation for safe methods
		if c.Request.Method == http.MethodGet ||
			c.Request.Method == http.MethodHead ||
			c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// Get CSRF token from cookie
		cookieToken, err := c.Cookie(auth.CSRFCookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CSRF token missing from cookie"})
			return
		}

		// Get CSRF token from header
		headerToken := c.GetHeader(auth.CSRFHeaderName)
		if headerToken == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CSRF token missing from header"})
			return
		}

		// Compare tokens (Double Submit Cookie pattern)
		if cookieToken != headerToken {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "CSRF token mismatch"})
			return
		}

		c.Next()
	}
}

// CSRFExempt creates a middleware that marks routes as exempt from CSRF validation
// Use this for routes that need to bypass CSRF (e.g., webhook endpoints)
func CSRFExempt() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("csrfExempt", true)
		c.Next()
	}
}
