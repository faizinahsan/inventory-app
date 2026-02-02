package http

import (
	"github.com/gin-gonic/gin"
	"inventory-app/internal/interfaces/handlers"
	"inventory-app/internal/interfaces/middleware"
)

// Router holds the HTTP router and handlers
type Router struct {
	engine         *gin.Engine
	productHandler *handlers.ProductHandler
}

// NewRouter creates a new HTTP router
func NewRouter(productHandler *handlers.ProductHandler) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// Add middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.CORS())

	return &Router{
		engine:         engine,
		productHandler: productHandler,
	}
}

// SetupRoutes sets up all HTTP routes
func (r *Router) SetupRoutes() {
	// Health check endpoint
	r.engine.GET("/health", r.healthCheck)

	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		// Product routes
		products := v1.Group("/products")
		{
			products.POST("", r.productHandler.CreateProduct)
			products.GET("", r.productHandler.ListProducts)
			products.GET("/search", r.productHandler.SearchProducts)
			products.GET("/low-stock", r.productHandler.GetLowStockProducts)
			products.GET("/:id", r.productHandler.GetProduct)
			products.PUT("/:id", r.productHandler.UpdateProduct)
			products.DELETE("/:id", r.productHandler.DeleteProduct)
		}

		// TODO: Add inventory routes
		// TODO: Add category routes
	}
}

// GetEngine returns the gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

// healthCheck handles health check requests
func (r *Router) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Inventory Management API is running",
	})
}
