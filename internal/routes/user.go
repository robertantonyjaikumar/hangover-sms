package routes

import (
	"sms/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	controller := new(handlers.UserRepo)
	router.POST("/", controller.Create)
	router.GET("/pagination", controller.GetAllByPagination)
	router.GET("/all", controller.GetAll)
	router.GET("/:id", controller.Get)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)

}
