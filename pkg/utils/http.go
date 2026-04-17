// Description: This file contains the helper functions for handling HTTP requests and responses.
// This file contains the helper functions for handling HTTP requests and responses.
package utils

import (
	"fmt"
	"net/http"
	"sms/internal/dtos/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, response.Response{Status: "success", Msg: cases.Title(language.English).String(message), Data: data})
}

func ErrorResponse(c *gin.Context, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, response.Response{Status: "fail", Msg: cases.Title(language.English).String(message), Data: data})
}
func UnAuthorizedResponse(c *gin.Context, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{Status: "fail", Msg: cases.Title(language.English).String(message), Data: data})
}

func GetPaginationParams(c *gin.Context) (offset int, limit int) {
	var err error
	offsetStr := c.Query("offset")
	offset, err = strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	limitStr := c.Query("limit")
	limit, err = strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return offset, limit
}

func GetHost(c *gin.Context) (host string) {
	origin := c.Request.Host
	scheme := "https"
	if c.Request.TLS == nil {
		scheme = "http"
	}

	host = fmt.Sprintf("%s://%s", scheme, origin)
	return
}
