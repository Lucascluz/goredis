package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas/goredis/internal/http"
)

const port = "8080"

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := router.Group("/api/v1")
	{
		api.GET("/items/:key", http.GET)
		api.POST("/items", http.SET)
		api.DELETE("/items/:key", http.DELETE)

		router.Run(":" + port)
	}
}
