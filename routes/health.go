package routes

import (
	"sms/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(router *gin.Engine) {
	health := new(controllers.HealthRepo)
	router.GET("/health", health.Status)
}
