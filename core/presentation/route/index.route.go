package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Tüm route'ları kaydederken health kontrolünü de ekleyin
func RegisterIndexRoutes(router *gin.Engine) {
	router.GET("/", HomeHandler)
	router.GET("/api/health", HealthCheck)
}
