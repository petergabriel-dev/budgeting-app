package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/petergabriel-dev/budgeting-app/internal/features/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthFlow(t *testing.T) {
	router, pool := SetupTestServer()
	defer pool.Close()

	// Clean up user if exists
	testEmail := "user_a@test.local"
	testPassword := "TestPass123!"

	_, err := pool.Exec(context.Background(), "DELETE FROM users WHERE email = $1", testEmail)
	assert.NoError(t, err)

	// 1. Register User
	t.Run("Register User", func(t *testing.T) {
		reqBody := auth.RegisterParams{
			Email:    testEmail,
			Password: testPassword,
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "User registered successfully", response["message"])

		userData := response["user"].(map[string]interface{})
		assert.Equal(t, testEmail, userData["email"])
	})

	// 2. Login User
	var sessionToken string
	t.Run("Login User", func(t *testing.T) {
		reqBody := auth.LoginParams{
			Email:    testEmail,
			Password: testPassword,
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Check Cookie
		cookies := w.Result().Cookies()
		var foundSession bool
		for _, cookie := range cookies {
			if cookie.Name == "session_token" {
				sessionToken = cookie.Value
				foundSession = true
				assert.True(t, cookie.HttpOnly)
			}
		}
		assert.True(t, foundSession, "Session cookie not found")
	})

	// 3. Access Protected Route (Me) with Cookie
	t.Run("Get Me (Protected)", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)

		// Add session cookie
		req.AddCookie(&http.Cookie{
			Name:  "session_token",
			Value: sessionToken,
		})

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		userData := response["user"].(map[string]interface{})
		assert.Equal(t, testEmail, userData["email"])
	})

	// 4. Logout
	t.Run("Logout", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/auth/logout", nil)
		req.AddCookie(&http.Cookie{
			Name:  "session_token",
			Value: sessionToken,
		})

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		cookies := w.Result().Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "session_token" {
				assert.Equal(t, "", cookie.Value, "Cookie should be cleared")
				assert.Equal(t, -1, cookie.MaxAge, "Cookie max age should be -1")
			}
		}
	})

	// 5. Access Protected Route without Cookie
	t.Run("Get Me (Unauthorized)", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/auth/me", nil)
		// No cookie

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
