package handlers

import (
	"fmt"
	"time"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"

	"sms/internal/dtos/request"
	"sms/internal/models"
	"sms/internal/repositories"
	"sms/pkg/utils"
)

type ProfileRepo struct{}

// ==================== GET PROFILE ====================

func (u ProfileRepo) Get(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := repositories.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", user)
}

// ==================== UPDATE PROFILE ====================

func (u ProfileRepo) Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Parse JSON fields
	//var updates map[string]interface{}

	if err := c.ShouldBind(&user); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	// Handle image upload if file exists
	file, err := c.FormFile("image")
	if err == nil {

		// Generate a unique file name (use the current timestamp and the original file extension)
		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("%v%s", id, ext)

		// Define the path where the image will be saved
		path := fmt.Sprintf("static/img/profile/%s", fileName)

		// Save the uploaded file to the defined path
		if err := c.SaveUploadedFile(file, path); err != nil {
			utils.ErrorResponse(c, "Failed to upload image", err.Error())
			return
		}

		// Add image URL to updates payload
		imageURL := fmt.Sprintf("%s/%s", utils.GetHost(c), path)
		user.Image = imageURL
	}

	// Update the user data in the DB
	if err := repositories.Update(c, models.User{}, user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	// Retrieve the updated user data
	if err := repositories.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	// Respond with success
	utils.SuccessResponse(c, "User updated successfully", user)
}

// ==================== FORGOT PASSWORD ====================

func (u ProfileRepo) ForgotPassword(c *gin.Context) {
	var input request.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	logger.Info("Forgot password requested", zap.String("email", input.Email))

	var user models.User
	if err := repositories.First(c, &user, "email = ?", input.Email); err != nil {
		utils.ErrorResponse(c, "Email not found", nil)
		return
	}

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

	resetLink := fmt.Sprintf("%s/v1/user-profile/reset-password?token=%s", utils.GetHost(c), accessToken)
	logger.Info("Password reset link generated", zap.String("link", resetLink))

	// need to fetch the email content from DB
	htmlBody := fmt.Sprintf(`
		<html>
			<body>
				<p>Hi %s,</p>
				<p>We received a request to reset your password. Click the link below to set a new one:</p>
				<p><a href="%s">Reset Your Password</a></p>
				<p>If you didn’t request this, you can safely ignore this email.</p>
				<br/>
				<p>Thanks,<br/>The Support Team</p>
			</body>
		</html>
	`, user.FirstName, resetLink)

	if err := utils.SendEmail(user.Email, "Reset Your Password", htmlBody); err != nil {
		utils.ErrorResponse(c, "Failed to send email", err.Error())
		return
	}

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

	var input request.ResetPasswordRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, "New password is required", nil)
		return
	}

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

	var user models.User
	if err := repositories.First(c, &user, "uuid = ?", userUUID); err != nil {
		utils.ErrorResponse(c, "User not found", nil)
		return
	}

	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		utils.ErrorResponse(c, "Failed to hash new password", nil)
		return
	}

	user.Password = hashedPassword
	if err := repositories.Save(c, &user); err != nil {
		utils.ErrorResponse(c, "Failed to update password", err.Error())
		return
	}

	utils.SuccessResponse(c, "Password reset successful", nil)
}
