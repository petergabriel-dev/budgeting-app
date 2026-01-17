---
name: create_go_feature
description: Scaffolds a new Go backend feature following the strict Package by Feature and Dependency Injection patterns.
---

# Create Go Feature

This skill guides you through creating a new feature module in the Go backend.
**Target Directory**: `src/backend/internal/features/<feature_name>/`

## 1. Prerequisites
- Confirm the feature name (e.g., `accounts`, `transactions`).
- Ensure `src/backend/internal/features/` exists.

## 2. File Structure Generation
You must create three files in the feature directory:

### A. `interface.go` (The Contract)
Defines the `Service` interface. This is what the Handler depends on.
```go
package <feature_name>

import "context"

type Service interface {
    // Define methods here, e.g.:
    // Create(ctx context.Context, userID string, params CreateParams) (*Model, error)
    // Get(ctx context.Context, id string) (*Model, error)
}
```

### B. `service.go` (The Implementation)
Implements the `Service` interface. Depends on `database.Queries` (sqlc).
```go
package <feature_name>

import (
    "github.com/yourusername/budgeting-app/internal/database"
)

type service struct {
    queries *database.Queries
}

// NewService is the factory function for dependency injection
func NewService(queries *database.Queries) Service {
    return &service{queries: queries}
}

// Implement interface methods here...
```

### C. `handler.go` (The HTTP Layer)
Depends ONLY on the `Service` interface.
```go
package <feature_name>

import (
    "github.com/gin-gonic/gin"
)

type Handler struct {
    service Service
}

// NewHandler is the factory function for dependency injection
func NewHandler(service Service) *Handler {
    return &Handler{service: service}
}

// Define handler methods here...
// func (h *Handler) Create(c *gin.Context) { ... }
```

## 3. Wiring (Manual Step)
Remind the user (or yourself) to update `cmd/server/main.go` to wire the new feature:
1.  Initialize Service: `svc := <feature>.NewService(queries)`
2.  Initialize Handler: `h := <feature>.NewHandler(svc)`
3.  Register Routes: `r.POST("/api/<feature>", h.Create)`
