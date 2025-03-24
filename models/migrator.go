package models

import (
	models "sms/models/utils"

	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
)

func GetTables() []interface{} {
	return []interface{}{
		&User{},
		&Todo{},
		// crud-generator-migration
	}
}

func MigrateDB() {
	var tables []interface{}
	tables = append(tables, GetTables()...)
	migrations := database.Migrations{
		DB:     database.Db,
		Models: tables,
	}
	database.RunMigrations(migrations)
}

func SeedDB() {
	seed := []models.Seed{
		// {Model: &[]Group{}, FileName: "user_groups.json", CreateFunc: SeedUserGroup},
		// {Model: &[]Role{}, FileName: "roles.json", CreateFunc: SeedRole},
		// {Model: &[]User{}, FileName: "users.json", CreateFunc: SeedUser},
		// {Model: &[]Todo{}, FileName: "todos.json", CreateFunc: SeedTodo},
		// crud-generator-seeds
	}
	for _, s := range seed {
		if err := models.SeedModel(s.FileName, s.Model, s.CreateFunc); err != nil {
			logger.Fatal("Error seeding roles", zap.Error(err))
		}
	}
}
