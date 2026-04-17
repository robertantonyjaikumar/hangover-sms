package models

import (
	"gorm.io/gorm"
	"time"
)

type PreModel struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy uint           `json:"created_by,omitempty"`
	UpdatedBy uint           `json:"updated_by,omitempty"`
	DeletedBy uint           `json:"deleted_by,omitempty"`
}

type PreModelUUIDAsID struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	UUID      string         `gorm:"uniqueIndex;default:gen_random_uuid();size:36" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy uint           `json:"created_by,omitempty"`
	UpdatedBy uint           `json:"updated_by,omitempty"`
	DeletedBy uint           `json:"deleted_by,omitempty"`
}

type PreModelWithUUID struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UUID      string         `gorm:"uniqueIndex;default:gen_random_uuid();size:36" json:"uuid"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy uint           `json:"created_by,omitempty"`
	UpdatedBy uint           `json:"updated_by,omitempty"`
	DeletedBy uint           `json:"deleted_by,omitempty"`
}
