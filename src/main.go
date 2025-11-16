// cmd/api/main.go
package main

import (
	"context"
	"fmt"
	"ginauth/src/config"
	"ginauth/src/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// --- injectables (overridable in tests) ---
var (
	loadEnv     = func() error { return godotenv.Load() }
	initDB      = config.InitDB
	setupRouter = routes.SetupRouter
	runHTTP     = func(r *gin.Engine, addr string) error { return r.Run(addr) }
	fatalf      = log.Fatalf
)

// split out the actual work so we can test
func run() error {
	if err := loadEnv(); err != nil {
		return fmt.Errorf("load env: %w", err)
	}
	initDB()
	// test DB connection
	conn := config.Conn()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.Ping(ctx); err != nil {
		return fmt.Errorf("database screening connection failed: %w", err)
	}
	fmt.Println("Database aml_screening connection successful!")

	r := setupRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	return runHTTP(r, ":"+port)
}

func main() {
	if err := run(); err != nil {
		fatalf("failed to start: %v", err)
	}
}
