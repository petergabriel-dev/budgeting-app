package auth

import (
	"context"

	"github.com/petergabriel-dev/budgeting-app/internal/database"
)

// AuthUser represents authenticated user data returned by the service
type AuthUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// RegisterParams contains the input for user registration
type RegisterParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginParams contains the input for user login
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResult contains the result of a successful login
type LoginResult struct {
	User         *AuthUser `json:"user"`
	SessionToken string    `json:"session_token"`
}

// Service defines the authentication service contract
type Service interface {
	// Register creates a new user account
	Register(ctx context.Context, params RegisterParams) (*AuthUser, error)

	// Login authenticates a user and creates a session
	Login(ctx context.Context, params LoginParams) (*LoginResult, error)

	// Logout invalidates a session token
	Logout(ctx context.Context, sessionToken string) error

	// ValidateSession validates a session token and returns the associated user
	ValidateSession(ctx context.Context, sessionToken string) (*AuthUser, error)
}

// Queries defines the database operations needed by the auth service
type Queries interface {
	CreateUser(ctx context.Context, arg database.CreateUserParams) (*database.User, error)
	GetUserByEmail(ctx context.Context, email string) (*database.User, error)
	CreateSession(ctx context.Context, arg database.CreateSessionParams) (*database.Session, error)
	GetSessionByToken(ctx context.Context, token string) (*database.GetSessionByTokenRow, error)
	DeleteSession(ctx context.Context, token string) error
}
