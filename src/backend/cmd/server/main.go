package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/petergabriel-dev/budgeting-app/internal/database"
	"github.com/petergabriel-dev/budgeting-app/internal/features/auth"
	"github.com/petergabriel-dev/budgeting-app/internal/shared/middleware"
)

func main() {
	// Load configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://budgeting:budgeting@localhost:5432/budgeting?sslmode=disable"
	}

	// Connect to database
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	// Verify database connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	log.Println("Connected to database")

	// Initialize database queries (sqlc)
	queries := database.New(pool)

	// Initialize services
	authService := auth.NewService(queries)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)

	// Initialize router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// Auth middleware for protected routes
		authMiddleware := middleware.SessionAuth(authService)

		// Register auth routes
		auth.RegisterRoutes(api, authHandler, authMiddleware)

		// Protected routes will be added here
		// Example:
		// protected := api.Group("")
		// protected.Use(authMiddleware)
		// protected.Use(middleware.CSRF())
		// {
		//     accounts.RegisterRoutes(protected, accountsHandler)
		// }
	}

	// Create server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server running on port %s", port)

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
