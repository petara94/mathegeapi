package models

import (
	"gorm.io/gorm"
)

type Task struct {
	Model
	MathEgeID     uint   `gorm:"unique;default:null" json:"math_ege_id"`
	TaskText      string `json:"task_text"`
	PatternTaskID *uint  `gorm:"default:null" json:"pattern_task_id"`

	PatternTask *PatternTask `json:"pattern_task"`
	Images      TaskImages   `json:"images"`
}

func (t *Task) BeforeUpdate(*gorm.DB) error {
	t.Model = Model{}
	return nil
}

func (t *Task) Allowed() map[string]interface{} {
	return map[string]interface{}{
		"math_ege_id":     t.MathEgeID,
		"task_text":       t.TaskText,
		"pattern_task_id": t.PatternTaskID,
	}
}

type Tasks []Task
