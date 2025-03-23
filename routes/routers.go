package routes

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"github.com/robertantonyjaikumar/hangover-common/middlewares"
)

func NewRouter() *gin.Engine {
	log := logger.GetZapLogger()
	router := gin.New()
	router.Use(middlewares.LogResponseAndRequestBodyMiddleware(log, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/health"},
	}))
	//router.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/health"}}))
	router.Use(gin.Recovery())
	HealthRoutes(router)
	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			AuthRoutes(authGroup)
		}
		todoGroup := v1.Group("todo")
		{
			TodoRoutes(todoGroup)
		}
	}
	return router
}
