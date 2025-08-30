package main

import (
	"log"

	"github.com/gin-gonic/gin"
	httpHandlers "github.com/lucas/goredis/internal/http"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Setup routes
	httpHandlers.SetupRoutes(r)

	// Start server
	log.Println("Starting GoRedis server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
