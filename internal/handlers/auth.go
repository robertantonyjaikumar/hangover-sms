package handlers

import (
	"sms/config"
	"sms/internal/dtos/request"
	"sms/internal/models"
	"sms/internal/repositories"
	"sms/pkg/utils"
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
//	@Param			LoginRequest	body		request.LoginRequest	true	"Add account"
//
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/login [post]
func (a AuthRepo) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := repositories.First(c, &user, "username = ?", req.Username); err != nil {
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
	user.Password = req.Password

	database.Db.Save(&user)

	utils.SuccessResponse(c, "", gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

// @Summary Refresh Token
// @Description Generate a new access token using the refresh token
// @Tags Auth
// @Accept json
// @Produce json
//
//	@Param			RefreshTokenRequest	body		request.RefreshTokenRequest	true	"Add account"
//
// @Success 200 Response body response.Response true "Success"
// @Failure 400 Response body response.Response true "Error"
// @Router /auth/refresh-token [post]
func (a AuthRepo) RefreshToken(c *gin.Context) {
	var req request.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := repositories.First(c, &user, "refresh_token = ?", req.RefreshToken); err != nil {
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
//	@Param			RefreshTokenRequest	body		request.RefreshTokenRequest	true	"Add account"
//
// @Success 200 Response body response.Response true "Success"
// @Failure 400 Response body response.Response true "Error"
// @Router /auth/logout [post]
func (a AuthRepo) Logout(c *gin.Context) {
	var req request.LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := repositories.First(c, &user, "refresh_token = ?", req.RefreshToken); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	user.RefreshToken = ""
	database.Db.Save(&user)

	utils.SuccessResponse(c, "", nil)
}
