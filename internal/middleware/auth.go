package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			// Fail safe: If no API key is configured, block everything to prevent accidental exposure
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server configuration error: API_KEY not set"})
			return
		}

		clientKey := c.GetHeader("X-API-KEY")

		if clientKey == "" || clientKey != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or missing API Key"})
			return
		}

		c.Next()
	}
}
