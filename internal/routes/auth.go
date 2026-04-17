package routes

import (
	"sms/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	auth := new(handlers.AuthRepo)
	router.POST("/login", auth.Login)
	// router.POST("/register", auth.Register)
	router.POST("/refresh", auth.RefreshToken)
	router.POST("/logout", auth.Logout)
	// router.POST("/logout-all", auth.LogoutAll)
	// router.POST("/forgot-password", auth.ForgotPassword)
	// router.POST("/reset-password", auth.ResetPassword)
	// router.POST("/verify-email", auth.VerifyEmail)
	// router.POST("/resend-verification-email", auth.ResendVerificationEmail)
	// router.POST("/change-password", auth.ChangePassword)
	// router.POST("/change-email", auth.ChangeEmail)
	// router.POST("/change-username", auth.ChangeUsername)
}
