package routes

import (
	"sms/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	auth := new(controllers.AuthRepo)
	router.POST("/login", auth.Login)
}
