This guide provides a production-ready architectural setup for a **Go (Chi)** backend and **React (Vite/TypeScript)** frontend. It uses a "Hybrid" development workflow: **Dual-Server** during development (for hot-reloading) and **Single Binary** for production (embedding React into Go).

### **Prerequisites**

* **Go** (v1.21+)
* **Node.js** (v20+ recommended)
* **VS Code** (recommended editor)

---

### **Phase 1: Project Scaffolding**

We will organize the project as a monorepo structure.

1. **Create the Root Directory**
```bash
mkdir go-react-app
cd go-react-app

```


2. **Initialize the Frontend (React/Vite)**
```bash
# Create React app in a 'web' directory
npm create vite@latest web -- --template react-ts

# Install dependencies
cd web
npm install
cd..

```


3. **Initialize the Backend (Go)**
```bash
# Initialize Go module
go mod init github.com/yourusername/go-react-app

# Install Chi router and Middleware
go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors

```



---

### **Phase 2: Frontend Configuration (Development Proxy)**

To avoid CORS errors during development, we configure Vite to proxy API requests to the Go backend.

1. **Open `web/vite.config.ts**` and update it:
```typescript
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Proxy /api requests to our Go server
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      }
    }
  },
  build: {
    // Build to the backend's static directory location (optional but cleaner)
    outDir: '../cmd/server/dist', 
    emptyOutDir: true, 
  }
})

```


2. **Test the Proxy Client:**
Edit `web/src/App.tsx` to fetch data from the Go API.
```tsx
import { useEffect, useState } from 'react'

function App() {
  const [message, setMessage] = useState('')

  useEffect(() => {
    fetch('/api/hello')
     .then(res => res.json())
     .then(data => setMessage(data.message))
  },)

  return (
    <div>
      <h1>Frontend: React + Vite</h1>
      <h2>Backend says: {message}</h2>
    </div>
  )
}

export default App

```



---

### **Phase 3: Backend Setup (The "SPA Handler")**

We need a robust Go server that serves API routes *and* handles the React routing fallback (so refreshing `/dashboard` doesn't 404).

1. **Create the Server Entry Point**
Create a file `main.go` (or `cmd/server/main.go` if you prefer standard layout).
```go
package main

import (
    "embed"
    "encoding/json"
    "io/fs"
    "log"
    "net/http"
    "strings"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

//go:embed web/dist/*
var staticFiles embed.FS

func main() {
    r := chi.NewRouter()

    // 1. Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Basic CORS for development (if needed explicitly)
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:string{"http://localhost:5173"},
        AllowedMethods:string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    }))

    // 2. API Routes (Grouped under /api)
    r.Route("/api", func(r chi.Router) {
        r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]string{
                "message": "Hello from Go!",
            })
        })
    })

    // 3. Serve React Frontend (Production Mode)
    // In Dev, we use Vite's proxy. In Prod, Go serves the files.
    // We strip "web/dist" because embed.FS includes the full path.
    contentStatic, _ := fs.Sub(staticFiles, "web/dist")
    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        // SPA Handler: If path is not API, try to serve static file.
        // If file doesn't exist, serve index.html (client-side routing).

        path := strings.TrimPrefix(r.URL.Path, "/")
        file, err := contentStatic.Open(path)
        if err!= nil {
            // File not found -> Serve index.html
            index, _ := contentStatic.ReadFile("index.html")
            w.Header().Set("Content-Type", "text/html")
            w.Write(index)
            return
        }
        // File exists -> Serve it
        file.Close()
        http.FileServer(http.FS(contentStatic)).ServeHTTP(w, r)
    })

    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", r)
}

```



---

### **Phase 4: Running the Application**

#### **Scenario A: Development Mode (Hot Reloading)**

You run both servers. Vite handles the UI and forwards API calls to Go.

1. **Terminal 1 (Backend):**
```bash
go run main.go
# Server listens on :8080

```


2. **Terminal 2 (Frontend):**
```bash
cd web
npm run dev
# Vite listens on :5173

```


3. **Visit:** `http://localhost:5173`
* You will see the React app.
* It will successfully fetch data from the Go API.
* Changes to Go code require a restart (or use `air`).
* Changes to React code hot-reload instantly.



#### **Scenario B: Production Build (Single Binary)**

You build the React app, embed it into Go, and ship one executable.

1. **Build React:**
```bash
cd web
npm run build
cd..

```


*This creates the `web/dist` folder.*
2. **Build Go Binary:**
```bash
go build -o myapp main.go

```


3. **Run Binary:**
```bash

```



./myapp
```
*   Visit `http://localhost:8080`.
*   Note: You are now accessing the Go server directly. It serves the HTML, CSS, JS, and the API.

### **Summary of Key Patterns Used**

1. **Vite Proxy:** Keeps frontend code clean. You write `fetch('/api/...')` rather than hardcoding `http://localhost:8080`.
2. **Go Embed:** Allows the final application to be a single binary file (easier deployment via Docker/Linux).
3. **SPA "NotFound" Handler:** The critical piece in `main.go`. It ensures that if a user refreshes the page at `/dashboard`, Go doesn't return a 404 but instead returns `index.html` so React Router can handle the view.