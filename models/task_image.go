package models

type TaskImage struct {
	Model
	FileName string `gorm:"not null"`
	FileType string `gorm:"not null"`
	TaskID   uint
}

type TaskImages []TaskImage
