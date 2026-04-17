package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthRepo struct{}

func (h HealthRepo) Status(c *gin.Context) {
	c.String(http.StatusOK, "Gin-starter is working")
}
