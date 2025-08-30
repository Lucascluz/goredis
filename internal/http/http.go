package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucas/goredis/internal/cache"
)

var cacheInstance = cache.NewCache()

type SetRequest struct {
	Value any    `json:"value"`
	TTL   *int64 `json:"ttl,omitempty"`
}

func SET(c *gin.Context) {
	// get key from header
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "key is required"})
		return
	}

	// get item from body
	var req SetRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
	}

	// insert pair on the map
	var err error
	if req.TTL != nil && *req.TTL > 0 {
		ttl := time.Duration(*req.TTL) * time.Second
		err = cacheInstance.SetWithTTL(key, req.Value, ttl)
	} else {
		err = cacheInstance.Set(key, req.Value)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to set value"})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"status": "succes",
		"key":    key,
	})
}

func GET(c *gin.Context) {
	// get key from header
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "key is required"})
		return
	}

	// check if key exists and its valid
	value, exists := cacheInstance.Get(key)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "key expired or do not exist"})
	}

	// return value
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"value":  value,
	})
}

func DELETE(c *gin.Context) {
	
}
