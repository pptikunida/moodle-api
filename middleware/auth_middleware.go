package middleware

import (
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/moodle-api/config"
	"github.com/rizkycahyono97/moodle-api/model/web"
	"net/http"
)

func ApiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requiredAPIKey := config.GetEnv("API_SECRET_KEY", "")
		if requiredAPIKey == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.ApiResponse{
				Code:    "500",
				Message: "Missing required API key",
			})
			return
		}

		//ambil api key
		suppliedAPIKey := c.GetHeader("X-API-Key")
		if suppliedAPIKey != requiredAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "401",
				Message: "Invalid API key",
			})
			return
		}

		//timing attack
		if subtle.ConstantTimeCompare([]byte(suppliedAPIKey), []byte(requiredAPIKey)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.ApiResponse{
				Code:    "401",
				Message: "Invalid API key",
			})
			return
		}

		c.Next()
	}
}
