package models

import (
	"fmt"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
)

type UserGroup struct {
	PreModel
	Name        string `gorm:"unique;not null"`
	Description string
	IsActive    *bool `json:"is_active"`
}

func (UserGroup) TableName() string {
	return "user_groups"
}

func SeedUserGroup(model interface{}) error {
	userGroups, ok := model.(*[]UserGroup)
	if !ok {
		return fmt.Errorf("invalid model type")
	}
	for _, userGroup := range *userGroups {
		if err := database.Db.FirstOrCreate(&userGroup, "name = ?", userGroup.Name).Error; err != nil {
			logger.Error("Error creating user seed: "+userGroup.Name, zap.Error(err))
			return err
		}
	}
	return nil
}
