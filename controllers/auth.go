package controllers

import (
	"sms/config"
	"sms/models"
	"sms/structs"
	"sms/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct{}

var jwtConfig = config.NewJwt()

// @Summary Login
// @Description Authenticate user and return access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
//
//	@Param			LoginRequest	body		structs.LoginRequest	true	"Add account"
//
// @Success 200 {object} structs.Response
// @Failure 400 {object} structs.Response
// @Router /auth/login [post]
func (a AuthRepo) Login(c *gin.Context) {
	var req structs.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := models.First(c, &user, "username = ?", req.Username); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	accessToken, _ := utils.GenerateToken(user.UUID, time.Duration(jwtConfig.AccessTokenExpireIn)*time.Second)
	refreshToken, _ := utils.GenerateToken(user.UUID, time.Duration(jwtConfig.RefreshTokenExpireIn)*time.Second)
	user.RefreshToken = refreshToken
	database.Db.Save(&user)

	utils.SuccessResponse(c, "", gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

// @Summary Refresh Token
// @Description Generate a new access token using the refresh token
// @Tags Auth
// @Accept json
// @Produce json
//
//	@Param			RefreshTokenRequest	body		structs.RefreshTokenRequest	true	"Add account"
//
// @Success 200 Response body structs.Response true "Success"
// @Failure 400 Response body structs.Response true "Error"
// @Router /auth/refresh-token [post]
func (a AuthRepo) RefreshToken(c *gin.Context) {
	var req structs.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := models.First(c, &user, "refresh_token = ?", req.RefreshToken); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	accessToken, _ := utils.GenerateToken(user.UUID, time.Duration(jwtConfig.AccessTokenExpireIn)*time.Second)
	utils.SuccessResponse(c, "", gin.H{"access_token": accessToken})
}

// @Summary Logout
// @Description Invalidate the refresh token
// @Tags Auth
// @Accept json
// @Produce json
// /
//
//	@Param			RefreshTokenRequest	body		structs.RefreshTokenRequest	true	"Add account"
//
// @Success 200 Response body structs.Response true "Success"
// @Failure 400 Response body structs.Response true "Error"
// @Router /auth/logout [post]
func (a AuthRepo) Logout(c *gin.Context) {
	var req structs.LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := models.First(c, &user, "refresh_token = ?", req.RefreshToken); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	user.RefreshToken = ""
	database.Db.Save(&user)

	utils.SuccessResponse(c, "", nil)
}
