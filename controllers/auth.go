package controllers

import (
	"sms/config"
	"sms/models"
	"sms/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct{}

var jwtConfig = config.NewJwt()

func (a AuthRepo) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

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

	accessToken, _ := utils.GenerateToken(user.ID, time.Duration(jwtConfig.AccessTokenExpireIn)*time.Second)
	refreshToken, _ := utils.GenerateToken(user.ID, time.Duration(jwtConfig.RefreshTokenExpireIn)*time.Second)
	user.RefreshToken = refreshToken
	database.Db.Save(&user)

	utils.SuccessResponse(c, "", gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

func (a AuthRepo) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	var user models.User
	if err := models.First(c, &user, "refresh_token = ?", req.RefreshToken); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	accessToken, _ := utils.GenerateToken(user.ID, time.Duration(jwtConfig.AccessTokenExpireIn)*time.Second)
	utils.SuccessResponse(c, "", gin.H{"access_token": accessToken})
}

func (a AuthRepo) Logout(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

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
