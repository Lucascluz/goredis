package http

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the cache API
func SetupRoutes(r *gin.Engine) {
	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Cache operations
		v1.POST("/keys/:key", SET)      // Set key-value with optional TTL
		v1.GET("/keys/:key", GET)       // Get value by key
		v1.DELETE("/keys/:key", DELETE) // Delete key
		v1.HEAD("/keys/:key", EXISTS)   // Check if key exists
		v1.GET("/keys", KEYS)           // List all keys

		// Administrative operations
		v1.POST("/flush", FLUSH) // Clear all data
	}

	// Health check endpoint
	r.GET("/health", HEALTH)

	// Info endpoint
	r.GET("/info", INFO)
}

func INFO(c *gin.Context) {
	c.JSON(200, gin.H{
		"server":  "GoRedis",
		"version": "1.0.0",
		"size":    cacheInstance.Size(),
	})
}
