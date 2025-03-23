package models

import (
	"fmt"

	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
)

type Todo struct {
	PreModelWithUUID
	Title       string `json:"title"`
	Slug        string `json:slug`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func SeedTodo(model interface{}) error {
	todos, ok := model.(*[]Todo)
	if !ok {
		return fmt.Errorf("invalid model type")
	}
	for _, todo := range *todos {
		if err := database.Db.FirstOrCreate(&todo, "slug = ?", todo.Slug).Error; err != nil {
			logger.Error("Error creating user seed: "+todo.Slug, zap.Error(err))
			return err
		}
	}
	return nil
}
