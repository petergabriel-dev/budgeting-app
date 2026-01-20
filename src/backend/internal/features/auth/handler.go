package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// Cookie names
	SessionCookieName = "session_token"
	CSRFCookieName    = "csrf_token"
	CSRFHeaderName    = "X-CSRF-Token"

	// Cookie settings
	CookiePath     = "/"
	CookieSecure   = false // Set to true in production with HTTPS
	CookieHTTPOnly = true
	CookieSameSite = http.SameSiteLaxMode
)

// Handler handles HTTP requests for authentication
type Handler struct {
	service Service
}

// NewHandler creates a new auth handler
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// RegisterRequest represents the register request body
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register handles user registration
// POST /api/v1/auth/register
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(c.Request.Context(), RegisterParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, ErrUserExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login handles user login
// POST /api/v1/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Login(c.Request.Context(), LoginParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	// Set session cookie (HttpOnly)
	c.SetSameSite(CookieSameSite)
	c.SetCookie(
		SessionCookieName,
		result.SessionToken,
		int(SessionDuration.Seconds()),
		CookiePath,
		"",
		CookieSecure,
		CookieHTTPOnly,
	)

	// Set CSRF token cookie (readable by JS for Double Submit Cookie pattern)
	csrfToken, _ := generateSessionToken()
	c.SetCookie(
		CSRFCookieName,
		csrfToken,
		int(SessionDuration.Seconds()),
		CookiePath,
		"",
		CookieSecure,
		false, // Not HttpOnly - needs to be readable by JavaScript
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    result.User,
	})
}

// Logout handles user logout
// POST /api/v1/auth/logout
func (h *Handler) Logout(c *gin.Context) {
	token, err := c.Cookie(SessionCookieName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Already logged out"})
		return
	}

	// Invalidate session in database
	if err := h.service.Logout(c.Request.Context(), token); err != nil {
		// Log error but continue - we still want to clear the cookie
	}

	// Clear cookies
	c.SetCookie(SessionCookieName, "", -1, CookiePath, "", CookieSecure, CookieHTTPOnly)
	c.SetCookie(CSRFCookieName, "", -1, CookiePath, "", CookieSecure, false)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Me returns the current authenticated user
// GET /api/v1/auth/me
func (h *Handler) Me(c *gin.Context) {
	// User should be set by auth middleware
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// RegisterRoutes registers auth routes on the given router group
func RegisterRoutes(rg *gin.RouterGroup, handler *Handler, authMiddleware gin.HandlerFunc) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.GET("/me", authMiddleware, handler.Me)
	}
}

// sessionDurationFromNow returns the session duration from now
func sessionDurationFromNow() time.Time {
	return time.Now().Add(SessionDuration)
}
