package controllers

import (
	"sms/models"
	"sms/utils"

	"github.com/gin-gonic/gin"
)

type TodoRepo struct{}

func (t TodoRepo) Create(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	err := models.Create(c, &todo)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", todo)
}

func (t TodoRepo) GetAll(c *gin.Context) {
	var todos []models.Todo

	if err := models.GetAll(c, &todos); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", todos)
}

func (t TodoRepo) Get(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := models.First(c, &todo, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", todo)
}

func (t TodoRepo) Update(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := models.Update(c, models.Todo{}, updates, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	if err := models.First(c, &todo, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", todo)
}

func (t TodoRepo) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := models.Delete(c, &models.Todo{}, id); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", nil)
}

func (t TodoRepo) GetAllByPagination(c *gin.Context) {
	var todos []models.Todo
	offset, limit := utils.GetPaginationParams(c)

	if err := models.GetAllByPagination(c, &todos, offset, limit); err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", todos)
}
