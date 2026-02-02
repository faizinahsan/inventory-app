package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

// CORS returns CORS middleware for Fiber
func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE",
	})
}

// JSONContentType ensures the request content type is JSON
func JSONContentType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "POST" || c.Method() == "PUT" || c.Method() == "PATCH" {
			contentType := c.Get("Content-Type")
			if contentType != "application/json" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Content-Type must be application/json",
				})
			}
		}
		return c.Next()
	}
}

// RequestID adds a request ID to each request
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Locals("RequestID", requestID)
		c.Set("X-Request-ID", requestID)
		return c.Next()
	}
}
