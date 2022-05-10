package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"reflect"
	"strings"
	"time"
)

type ID uint64

type Model struct {
	ID        ID             `gorm:"primarykey" json:"id" entity:"@id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;autoDeleteTime:milli" json:"-"`
}

func AllowedToUpdate(entity any) map[string]any {
	t := reflect.TypeOf(entity).Elem()
	v := reflect.ValueOf(entity).Elem()
	elems := make(map[string]any)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			continue
		}
		if !utils.Contains(strings.Split(f.Tag.Get("entity"), ";"), "protected") {
			elems[f.Name] = v.FieldByName(f.Name).Interface()
		}
	}
	return elems
}

// GetIdFieldNameOfEntity возвращает имя поля ключа
func GetIdFieldNameOfEntity(entity any) (id string) {
	t := reflect.TypeOf(entity).Elem()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			id = GetIdFieldNameOfEntity(reflect.New(f.Type).Interface())
			if id != "" {
				return
			}
		}
		for _, j := range strings.Split(f.Tag.Get("entity"), ";") {
			if j == "@id" {
				id = f.Name
				return
			}
		}
	}
	return
}
