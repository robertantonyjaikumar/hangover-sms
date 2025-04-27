package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
	"sms/models"
	"sms/structs"
	"sms/utils"
	"time"
)

type ProfileRepo struct{}

// ==================== GET PROFILE ====================

func (u ProfileRepo) Get(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", user)
}

// ==================== UPDATE PROFILE ====================

func (u ProfileRepo) Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	var updates map[string]interface{}

	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := models.Update(c, models.User{}, updates, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := models.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", user)
}

// ==================== FORGOT PASSWORD ====================

func (u ProfileRepo) ForgotPassword(c *gin.Context) {
	var input structs.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	logger.Info("Forgot password requested", zap.Any("email", input.Email))

	var user models.User
	if err := models.First(c, &user, "email = ?", input.Email); err != nil {
		utils.ErrorResponse(c, "Email not found", nil)
		return
	}

	// Generate tokens
	accessToken, err := utils.GenerateToken(user.UUID, time.Duration(jwtConfig.AccessTokenExpireIn)*time.Second)
	if err != nil {
		utils.ErrorResponse(c, "Failed to generate access token", nil)
		return
	}

	refreshToken, err := utils.GenerateToken(user.UUID, time.Duration(jwtConfig.RefreshTokenExpireIn)*time.Second)
	if err != nil {
		utils.ErrorResponse(c, "Failed to generate refresh token", nil)
		return
	}

	user.RefreshToken = refreshToken
	if err := database.Db.Save(&user).Error; err != nil {
		utils.ErrorResponse(c, "Failed to save user refresh token", nil)
		return
	}

	// Build dynamic reset link
	resetLink := fmt.Sprintf("%s/v1/user-profile/reset-password?token=%s", utils.GetHost(c), accessToken)
	logger.Info("Password reset link generated", zap.String("link", resetLink))

	// TODO: send reset link via email

	utils.SuccessResponse(c, "Password reset link sent to email", gin.H{
		"reset_link": resetLink,
	})
}

// ==================== RESET PASSWORD ====================

func (u ProfileRepo) ResetPassword(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		utils.ErrorResponse(c, "Reset token is required", nil)
		return
	}

	var input structs.ResetPasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, "New password is required", nil)
		return
	}

	// Parse token
	claims, err := utils.ParseToken(token)
	if err != nil {
		utils.UnAuthorizedResponse(c, "Invalid or expired reset token", nil)
		return
	}

	userUUID, ok := (*claims)["user_id"].(string)
	if !ok || userUUID == "" {
		utils.UnAuthorizedResponse(c, "Invalid or missing token subject", nil)
		return
	}

	// Fetch user
	var user models.User
	if err := models.First(c, &user, "uuid = ?", userUUID); err != nil {
		utils.ErrorResponse(c, "User not found", nil)
		return
	}

	//// Hash password
	//hashedPassword, err := utils.HashPassword(input.NewPassword)
	//if err != nil {
	//	utils.ErrorResponse(c, "Failed to hash new password", nil)
	//	return
	//}

	// Save new password
	user.Password = input.NewPassword
	if err := models.Save(c, &user); err != nil {
		utils.ErrorResponse(c, "Failed to update password", err.Error())
		return
	}

	utils.SuccessResponse(c, "Password reset successful", nil)
}
