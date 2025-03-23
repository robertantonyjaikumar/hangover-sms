package routes

import (
	"hangover/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.RouterGroup) {
	controller := new(controllers.TodoRepo)
	router.POST("/", controller.Create)
	router.GET("/pagination", controller.GetAllByPagination)
	router.GET("/all", controller.GetAll)
	router.GET("/:id", controller.Get)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)

}
