package models

import (
	"fmt"
	"time"

	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	PreModelWithUUID
	Username      string     `json:"username" gorm:"uniqueIndex;not null"`
	FirstName     string     `json:"first_name"`
	MiddleName    string     `json:"middle_name"`
	LastName      string     `json:"last_name"`
	DisplayName   string     `json:"display_name"`
	DOB           *time.Time `json:"dob"`
	ContactNumber string     `json:"contact_number"`
	CountryCode   string     `json:"country_code"`
	Email         string     `json:"email" gorm:"uniqueIndex;not null"`
	Password      string     `json:"password"`
	RefreshToken  string     `gorm:"type:text"`
	LastLoginAt   *time.Time `json:"last_login_at"`
	LastLoginIP   string     `json:"last_login_ip"`
}

func SeedUser(model interface{}) error {
	users, ok := model.(*[]User)
	if !ok {
		return fmt.Errorf("invalid model type")
	}
	for _, user := range *users {
		if err := database.Db.FirstOrCreate(&user, "username = ?", user.Username).Error; err != nil {
			logger.Error("Error creating user seed: "+user.Username, zap.Error(err))
			return err
		}
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if hashErr != nil {
			logger.Error("Error hashing password", zap.Error(hashErr))
			return hashErr
		}
		u.Password = string(hashedPassword)
		if err != nil {
			logger.Error("Error hashing password", zap.Error(err))
			return err
		}
	}
	middleName := ""
	if u.MiddleName != "" {
		middleName = fmt.Sprintf(" %s", u.MiddleName)
	}
	u.DisplayName = fmt.Sprintf("%s%s %s", u.FirstName, middleName, u.LastName)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if hashErr != nil {
			logger.Error("Error hashing password", zap.Error(hashErr))
			return hashErr
		}
		u.Password = string(hashedPassword)
		if err != nil {
			logger.Error("Error hashing password", zap.Error(err))
			return err
		}
	}
	middleName := ""
	if u.MiddleName != "" {
		middleName = fmt.Sprintf(" %s", u.MiddleName)
	}
	u.DisplayName = fmt.Sprintf("%s%s %s", u.FirstName, middleName, u.LastName)

	logger.Info("BeforeUpdate", zap.String("DisplayName", u.DisplayName))
	return nil
}
