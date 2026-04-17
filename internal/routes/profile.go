package routes

import (
	"sms/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.RouterGroup) {
	controller := new(handlers.ProfileRepo)

	router.GET("/:id", controller.Get)

	router.POST("/forgot-password", controller.ForgotPassword)
	router.POST("/reset-password", controller.ResetPassword)

	router.PUT("/:id", controller.Update)

}
