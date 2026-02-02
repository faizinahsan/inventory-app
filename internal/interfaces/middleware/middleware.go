package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS adds CORS headers to responses
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// JSONContentType ensures the request content type is JSON
func JSONContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.GetHeader("Content-Type")
			if contentType != "application/json" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// RequestID adds a request ID to each request
func RequestID() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			// Generate a simple request ID if not provided
			requestID = generateRequestID()
		}

		c.Set("RequestID", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	})
}

// generateRequestID generates a simple request ID
func generateRequestID() string {
	// Simple implementation - in production, use a proper UUID or similar
	return "req-" + randomString(8)
}

// randomString generates a random string of specified length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[len(charset)/2] // Simplified for demo
	}
	return string(b)
}
