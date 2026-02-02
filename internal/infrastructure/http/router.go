package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"inventory-app/internal/interfaces/handlers"
	"inventory-app/internal/interfaces/middleware"
)

// Router holds the HTTP router and handlers
type Router struct {
	app            *fiber.App
	productHandler *handlers.ProductHandler
}

// NewRouter creates a new HTTP router
func NewRouter(productHandler *handlers.ProductHandler) *Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(middleware.CORS())

	return &Router{
		app:            app,
		productHandler: productHandler,
	}
}

// SetupRoutes sets up all HTTP routes
func (r *Router) SetupRoutes() {
	// Health check endpoint
	r.app.Get("/health", r.healthCheck)

	// API v1 routes
	v1 := r.app.Group("/api/v1")
	{
		// Product routes
		products := v1.Group("/products")
		{
			products.Post("/", r.productHandler.CreateProduct)
			products.Get("/", r.productHandler.ListProducts)
			products.Get("/search", r.productHandler.SearchProducts)
			products.Get("/low-stock", r.productHandler.GetLowStockProducts)
			products.Get("/:id", r.productHandler.GetProduct)
			products.Put("/:id", r.productHandler.UpdateProduct)
			products.Delete("/:id", r.productHandler.DeleteProduct)
		}

		// TODO: Add inventory routes
		// TODO: Add category routes
	}
}

// GetApp returns the fiber app
func (r *Router) GetApp() *fiber.App {
	return r.app
}

// healthCheck handles health check requests
func (r *Router) healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Inventory Management API is running",
	})
}
