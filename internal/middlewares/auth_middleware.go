package middlewares

import (
	"sms/pkg/utils"

	"github.com/gin-gonic/gin"
)

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}
	return ""
}

func extractUserID(claims map[string]interface{}) (string, bool) {
	userID, ok := claims["user_id"].(string)
	return userID, ok
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			utils.UnAuthorizedResponse(c, "Authorization header is required", nil)
			return
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.UnAuthorizedResponse(c, err.Error(), nil)
			return
		}
		c.Set("user_id", (*claims)["user_id"])
		c.Next()

		// if userID, ok := extractUserID(map[string]interface{}(*claims)); ok {
		// 	c.Set("user_id", userID)
		// 	c.Next()
		// } else {
		// 	utils.UnAuthorizedResponse(c, "Invalid token claims", nil)
		// }
	}
}
