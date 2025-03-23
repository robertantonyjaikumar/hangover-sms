package models

import (
	"fmt"
	models "hangover/models/utils"

	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
)

type User struct {
	PreModelWithUUID
	Username     string `json:"username" gorm:"uniqueIndex;not null"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name"`
	LastName     string `json:"last_name"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"password_hash"`
	UserGroup    uint   `json:"user_group"`
	Group        Group  `gorm:"foreignKey:UserGroup"`
	Roles        []Role `gorm:"many2many:user_roles;"`
	IsActive     *bool  `json:"is_active"`
}

func (u *User) TableName() string {
	return "users"
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	if err := database.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ValidateUserByUserNameAndPassword(username, password string) (*User, error) {
	user, err := GetUserByUserName(username)
	if err != nil {
		return nil, err
	}
	if models.ValidatePassword(user.PasswordHash, password) {
		return user, nil
	}
	return nil, fmt.Errorf("invalid username or password")
}

func SeedUser(model interface{}) error {
	users, ok := model.(*[]User)
	if !ok {
		return fmt.Errorf("invalid model type")
	}
	for _, user := range *users {
		user.PasswordHash, _ = models.HashPassword(user.PasswordHash)
		if err := database.Db.FirstOrCreate(&user, "username = ?", user.Username).Error; err != nil {
			logger.Error("Error creating user seed: "+user.Username, zap.Error(err))
			return err
		}
	}
	return nil
}
