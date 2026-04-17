package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID           uint           `gorm:"primaryKey"`
	AccessToken2 string         `json:"name" gorm:"unique;not null"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (t *Token) TableName() string {
	return "tokens"
}
