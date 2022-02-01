package models

type Task struct {
	Model
	MathEgeID     uint   `gorm:"unique"`
	TaskText      string `json:"task_text"`
	PatternTaskID uint   `json:"pattern_task_id"`

	PatternTask PatternTask `json:"pattern_task"`
	Images      []TaskImage `json:"images"`
}