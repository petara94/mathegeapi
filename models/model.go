package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `json:"updated_at;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index;autoDeleteTime:milli" json:"-"`
}
