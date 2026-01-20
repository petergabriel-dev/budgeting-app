package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/petergabriel-dev/budgeting-app/internal/database"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidSession     = errors.New("invalid or expired session")
)

const (
	// SessionDuration defines how long a session is valid
	SessionDuration = 7 * 24 * time.Hour // 7 days
	// BcryptCost defines the bcrypt hashing cost
	BcryptCost = 12
)

type service struct {
	queries *database.Queries
}

// NewService creates a new auth service instance
func NewService(queries *database.Queries) Service {
	return &service{queries: queries}
}

// Register creates a new user with hashed password
func (s *service) Register(ctx context.Context, params RegisterParams) (*AuthUser, error) {
	// Check if user already exists
	_, err := s.queries.GetUserByEmail(ctx, params.Email)
	if err == nil {
		return nil, ErrUserExists
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), BcryptCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user, err := s.queries.CreateUser(ctx, database.CreateUserParams{
		Email:        params.Email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return &AuthUser{
		ID:    formatUUID(user.ID),
		Email: user.Email,
	}, nil
}

// Login authenticates a user and creates a session
func (s *service) Login(ctx context.Context, params LoginParams) (*LoginResult, error) {
	// Get user by email
	user, err := s.queries.GetUserByEmail(ctx, params.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate session token
	token, err := generateSessionToken()
	if err != nil {
		return nil, err
	}

	// Create session
	expiresAt := time.Now().Add(SessionDuration)
	_, err = s.queries.CreateSession(ctx, database.CreateSessionParams{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		User: &AuthUser{
			ID:    formatUUID(user.ID),
			Email: user.Email,
		},
		SessionToken: token,
	}, nil
}

// Logout invalidates a session
func (s *service) Logout(ctx context.Context, sessionToken string) error {
	return s.queries.DeleteSession(ctx, sessionToken)
}

// ValidateSession validates a session token and returns the user
func (s *service) ValidateSession(ctx context.Context, sessionToken string) (*AuthUser, error) {
	session, err := s.queries.GetSessionByToken(ctx, sessionToken)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidSession
		}
		return nil, err
	}

	return &AuthUser{
		ID:    formatUUID(session.UserID_2),
		Email: session.UserEmail,
	}, nil
}

// generateSessionToken creates a cryptographically secure random token
func generateSessionToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// formatUUID converts a pgtype.UUID to a string
func formatUUID(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	return hex.EncodeToString(u.Bytes[0:4]) + "-" +
		hex.EncodeToString(u.Bytes[4:6]) + "-" +
		hex.EncodeToString(u.Bytes[6:8]) + "-" +
		hex.EncodeToString(u.Bytes[8:10]) + "-" +
		hex.EncodeToString(u.Bytes[10:16])
}
