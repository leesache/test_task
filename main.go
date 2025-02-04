package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/handlers"
	"example.com/middleware"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	models.InitDB()

	// Apply migrations
	models.RunMigrations()

	// Create a new Gin router
	router := gin.Default()

	// Add the error handling middleware
	router.Use(middleware.ErrorHandlerMiddleware)

	// Public routes
	router.POST("/users/register", handlers.RegisterUser)

	// Protected routes (require authentication)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users/leaderboard", handlers.GetUsers)
		protected.GET("/users/:id/status", handlers.GetUserInfo)
		protected.POST("/users/:id/referrer", handlers.ReferUser)
		protected.POST("/users/:id/task/complete", handlers.TaskComplete)
	}

	// Create an HTTP server
	server := &http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown logic
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Block until a signal is received

	log.Println("Shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
