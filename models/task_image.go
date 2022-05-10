package models

type TaskImage struct {
	Model
	FileName string `gorm:"not null" json:"file_name"`
	FileType string `gorm:"not null" json:"file_type"`
	TaskID   uint   `json:"task_id"`
}

type TaskImages []TaskImage
