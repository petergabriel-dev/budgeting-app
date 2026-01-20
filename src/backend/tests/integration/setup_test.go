package integration

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/petergabriel-dev/budgeting-app/internal/database"
	"github.com/petergabriel-dev/budgeting-app/internal/features/auth"
	"github.com/petergabriel-dev/budgeting-app/internal/shared/middleware"
)

// SetupTestServer initializes the database connection and router for testing
func SetupTestServer() (*gin.Engine, *pgxpool.Pool) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://budgeting:budgeting@localhost:5432/budgeting?sslmode=disable"
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Verify database connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	queries := database.New(pool)
	authService := auth.NewService(queries)
	authHandler := auth.NewHandler(authService)

	gin.SetMode(gin.TestMode)
	r := gin.New() // Use New() to avoid default logger/recovery logs during tests
	r.Use(gin.Recovery())

	// CORS configuration (match main.go)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api/v1")
	{
		authMiddleware := middleware.SessionAuth(authService)
		auth.RegisterRoutes(api, authHandler, authMiddleware)
	}

	return r, pool
}
