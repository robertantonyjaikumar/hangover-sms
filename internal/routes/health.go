package routes

import (
	"sms/internal/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(router *gin.Engine) {
	health := new(handlers.HealthRepo)
	router.GET("/health", health.Status)
}
