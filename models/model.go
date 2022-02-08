package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;autoDeleteTime:milli" json:"-"`
}

//AllowedToUpdate Обязует замещать данные которые
//защищены от обновления перед транзакией
type AllowedToUpdate interface {
	callbacks.BeforeUpdateInterface

	//Allowed Возвращает поля и их значения для изменения
	Allowed() map[string]interface{}
}
