package controllers

import (
	"hangover/models"
	"hangover/utils"

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

	todos, err := models.GetAll(c)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", todos)
}

func (t TodoRepo) Get(c *gin.Context) {
	id := c.Param("id")
	todo, err := models.First(c, id)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, "", todo)
}

func (t TodoRepo) Update(c *gin.Context) {
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
	todo, err := models.First(c, id)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", todo)
}

func (t TodoRepo) Delete(c *gin.Context) {
	id := c.Param("id")
	err := models.Delete(c, &models.Todo{}, id)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", nil)
}

func (t TodoRepo) GetAllByPagination(c *gin.Context) {
	offset, limit := utils.GetPaginationParams(c)
	todos, err := models.GetAllByPagination(c, offset, limit)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), nil)
		return
	}
	utils.SuccessResponse(c, "", todos)
}
