package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"inventory-app/internal/application/usecases"
	"inventory-app/internal/domain/services"
	"inventory-app/internal/infrastructure/config"
	"inventory-app/internal/infrastructure/database"
	"inventory-app/internal/infrastructure/database/postgres"
	httpInfra "inventory-app/internal/infrastructure/http"
	"inventory-app/internal/interfaces/handlers"
	"inventory-app/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	appLogger, err := logger.NewLogger(cfg.Logger.Level, cfg.Logger.Format)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer appLogger.Sync()

	appLogger.Info("Starting Inventory Management Application")

	// Initialize database connection
	db, err := database.NewConnection(cfg)
	if err != nil {
		appLogger.Fatal("Failed to connect to database", appLogger.WithField("error", err))
	}
	defer db.Close()

	appLogger.Info("Database connection established")

	// Initialize repositories
	productRepo := postgres.NewProductRepository(db)

	// Initialize services
	inventoryService := services.NewInventoryService(productRepo, nil) // Transaction repo will be added later

	// Initialize use cases
	productUseCase := usecases.NewProductUseCase(productRepo, nil, inventoryService) // Category repo will be added later

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productUseCase)

	// Initialize HTTP router
	router := httpInfra.NewRouter(productHandler)
	router.SetupRoutes()

	// Create HTTP server
	server := &http.Server{
		Addr:           cfg.ServerAddress(),
		Handler:        router.GetEngine(),
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// Start server in a goroutine
	go func() {
		appLogger.Info("Starting HTTP server", appLogger.WithField("address", cfg.ServerAddress()))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatal("Failed to start server", appLogger.WithField("error", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		appLogger.Fatal("Server forced to shutdown", appLogger.WithField("error", err))
	}

	appLogger.Info("Server exited")
}
