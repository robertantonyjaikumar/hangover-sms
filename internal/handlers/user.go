package handlers

import (
	"sms/internal/models"
	"sms/internal/repositories"
	"sms/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserRepo struct{}

func (u UserRepo) Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	err := repositories.Create(c, &user)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", user)
}

func (u UserRepo) GetAll(c *gin.Context) {
	var users []models.User

	if err := repositories.GetAll(c, &users); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", users)
}

func (u UserRepo) Get(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := repositories.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", user)
}

func (u UserRepo) Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := repositories.Update(c, models.User{}, updates, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := repositories.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", user)
}

func (u UserRepo) Delete(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := repositories.First(c, &user, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := repositories.Delete(c, &user); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", nil)
}

func (u UserRepo) GetAllByPagination(c *gin.Context) {
	var users []models.User
	offset, limit := utils.GetPaginationParams(c)

	if err := repositories.GetAllByPagination(c, &users, offset, limit); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", users)
}
